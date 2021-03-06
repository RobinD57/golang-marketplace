package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/RobinD57/golang-marketplace/pkg/chat"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
var redisHost string
var redisPassword string
var validate *validator.Validate

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,string"`
	PublicAddress string             `bson:"publicAddress,omitempty" json:"publicAddress,omitempty"`
	Nonce         string             `bson:"nonce,omitempty" json:"nonce,omitempty"` // keep it as string, could be big rnd int
	Reviews       []Review           `bson:"reviews,omitempty" json:"reviews"`
}

func (u *User) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("ethaddress", validateAddress)
	return validate.Struct(u)
}

func validateAddress(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)
	matches := re.FindAllString(fl.Field().String(), -1) // slice of strings
	if len(matches) != 1 {
		return false
	}
	return true
}

type Review struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,string"`
	UserPubAddress string             `bson:"userPubAddress" json:"userPubAddress" validate:"required"` // the one getting the review
	Rating         int                `bson:"rating,omitempty" json:"rating, string, omitempty" validate:"required,gte=1,lte=5"`
	Content        string             `bson:"content,omitempty" json:"content,omitempty"`
	Reviewer       string             `bson:"reviewer, omitempty" json:"reviewer"`
}

func (r *Review) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type Listing struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,string"`
	Name        string             `bson:"name,omitempty" json:"name" validate:"required"`
	Seller      string             `bson:"seller,omitempty" json:"seller,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Price       float64            `bson:"price,omitempty" json:"price, string" validate:"required"`
	Photos      []string           `bson:"photos,omitempty" json:"photos" validate:"required"` // Cloudinary?
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
}

func (l *Listing) Validate() error {
	validate := validator.New()
	return validate.Struct(l)
}

type Transaction struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,string"`
	Listing primitive.ObjectID `bson:"listing" json:"listing,string"`
	Buyer   string             `bson:"buyer,omitempty" json:"buyer,omitempty"`
	Seller  string             `bson:"seller,omitempty" json:"seller,omitempty"`
	Status  string             `bson:"status,omitempty" json:"status,omitempty" validate:"status"`
}

func (t *Transaction) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("status", validateStatus)
	return validate.Struct(t)
}

func validateStatus(fl validator.FieldLevel) bool {
	return fl.Field().String() == "open" || fl.Field().String() == "escrow" || fl.Field().String() == "done"
}

type Connection struct {
	Listings     *mongo.Collection
	Transactions *mongo.Collection
	Reviews      *mongo.Collection
	Users        *mongo.Collection
}

func (connection Connection) CreateListingEndpoint(w http.ResponseWriter, r *http.Request) { // no proper validations for now
	var listing Listing
	if err := json.NewDecoder(r.Body).Decode(&listing); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err := listing.Validate(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	listing.CreatedAt = time.Now()
	result, err := connection.Listings.InsertOne(ctx, listing)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (connection Connection) GetListingsEndpoint(w http.ResponseWriter, r *http.Request) {
	var listings []Listing
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := connection.Listings.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &listings); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(listings)
}

func (connection Connection) GetListingEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var listing Listing
	var result bson.M
	json.NewDecoder(r.Body).Decode(&listing)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := connection.Listings.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (connection Connection) UpdateListingEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var listing Listing
	json.NewDecoder(r.Body).Decode(&listing)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := connection.Listings.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", listing},
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (connection Connection) DeleteListingEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := connection.Listings.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (connection Connection) GetTransactionsEndpoint(w http.ResponseWriter, r *http.Request) {
	var transactions []bson.M
	params := mux.Vars(r)
	publicAddress, _ := params["publicAddress"]
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := connection.Transactions.Find(ctx, bson.M{"$or": []interface{}{
		bson.M{"seller": publicAddress},
		bson.M{"buyer": publicAddress},
	},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &transactions); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(transactions)
}

func (connection Connection) CreateTransactionEndpoint(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	transaction.Listing = id
	transaction.Status = "open"
	result, err := connection.Transactions.InsertOne(ctx, transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (connection Connection) DeleteTransactionEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := connection.Transactions.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (connection Connection) FindOrCreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	params := mux.Vars(r)
	publicAddress, _ := params["publicAddress"]
	var user User
	var reviews []bson.M
	var result bson.M
	json.NewDecoder(r.Body).Decode(&user)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := connection.Reviews.Find(ctx, bson.D{{"userPubAddress", publicAddress}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &reviews); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err := user.Validate(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	filter := bson.D{{"publicAddress", publicAddress}}
	update := bson.D{{"$set", bson.D{{"Reviews", reviews}}}}
	err_ := connection.Users.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err_ != nil {
		if err_ == mongo.ErrNoDocuments {
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err_.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (connection Connection) CreateReviewEndpoint(w http.ResponseWriter, r *http.Request) {
	var review Review
	params := mux.Vars(r)
	publicAddress, _ := params["publicAddress"]
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	review.UserPubAddress = publicAddress
	if err := review.Validate(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	result, err := connection.Reviews.InsertOne(ctx, review)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	params := mux.Vars(r)
	user, _ := params["username"]
	peer, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("websocket conn failed", err)
	}
	chatSession := chat.NewChatSession(user, peer)
	chatSession.Start()
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func init() {
	redisHost = goDotEnvVariable("REDIS_HOST")
	if redisHost == "" {
		log.Fatal("missing REDIS_HOST env var")
	}

	redisPassword = goDotEnvVariable("REDIS_PASSWORD")
	if redisPassword == "" {
		log.Fatal("missing REDIS_PASSWORD env var")
	}
}

func main() {
	dotenv := goDotEnvVariable("MONGO_ATLAS")

	// connect to DB
	client, err := mongo.NewClient(options.Client().ApplyURI(dotenv))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	marketplaceDatabase := client.Database("marketplace")
	listingsCollection := marketplaceDatabase.Collection("listings")
	transactionsCollection := marketplaceDatabase.Collection("transactions")
	usersCollection := marketplaceDatabase.Collection("users")
	reviewsCollection := marketplaceDatabase.Collection("reviews")

	connection := Connection{Listings: listingsCollection, Transactions: transactionsCollection, Users: usersCollection, Reviews: reviewsCollection}

	router := mux.NewRouter()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/listing", connection.CreateListingEndpoint).Methods("POST", "OPTIONS")
	router.HandleFunc("/listings", connection.GetListingsEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/listing/{id}", connection.GetListingEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/listing/{id}", connection.UpdateListingEndpoint).Methods("PUT", "OPTIONS")
	router.HandleFunc("/listing/{id}", connection.DeleteListingEndpoint).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/listing/{id}/transaction", connection.CreateTransactionEndpoint).Methods("POST", "OPTIONS")
	router.HandleFunc("/transaction/{id}", connection.DeleteTransactionEndpoint).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/user/{publicAddress}", connection.FindOrCreateUserEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/chat/{username}", websocketHandler)
	router.HandleFunc("/user/{publicAddress}/review", connection.CreateReviewEndpoint).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/{publicAddress}/transactions", connection.GetTransactionsEndpoint).Methods("GET", "OPTIONS")
	http.ListenAndServe(":8080", handlers.CORS(header, methods, origins)(router))

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	log.Println("exit signalled")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	chat.Cleanup()
	log.Println("chat app exited")
}
