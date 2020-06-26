package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

var tpl = template.Must(template.ParseFiles("static/index.html"))

// need to add validations to structs!!

type User struct {
	ID            primitive.ObjectID `bson:"_id" json:"id,string"`
	PublicAddress string             `bson:"publicAddress,omitempty" json:"publicAddress,omitempty"`
	Nonce         string             `bson:"nonce,omitempty" json:"nonce,omitempty"` // keep it as string, could be big rnd int
	Reviews       []Review           `bson:"reviews,omitempty" json:"reviews,string"`
}

type Review struct {
	ID      primitive.ObjectID `bson:"_id" json:"id,string"`
	Rating  int                `bson:"rating,omitempty" json:"rating, string, omitempty"` // range missing
	Content string             `bson:"content,omitempty" json:"content,omitempty"`
}

type Listing struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,string"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Seller      primitive.ObjectID `bson:"seller,omitempty" json:"seller,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Price       float64            `bson:"price,omitempty" json:"price, string, omitempty"`
	Photo       string             `bson:"photo,omitempty" json:"photo,omitempty"` // Cloudinary?
}

type Purchase struct {
	ID      primitive.ObjectID `bson:"_id" json:"id,string"`
	Listing primitive.ObjectID `bson:"listing,omitempty" json:"listing,string"`
	Buyer   primitive.ObjectID `bson:"buyer,omitempty" json:"buyer,omitempty"`
	Seller  primitive.ObjectID `bson:"seller,omitempty" json:"seller,omitempty"`
}

type Connection struct {
	Listings  *mongo.Collection
	Purchases *mongo.Collection
	Reviews   *mongo.Collection
	Users     *mongo.Collection
}

func (connection Connection) CreateListingEndpoint(w http.ResponseWriter, r *http.Request) { // no proper validations for now
	w.Header().Set("content-type", "application/json")
	var listing Listing
	if err := json.NewDecoder(r.Body).Decode(&listing); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
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
	w.Header().Set("content-type", "application/json")
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
	w.Header().Set("content-type", "application/json")
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

func (connection Connection) CreatePurchaseEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var purchase Purchase
	if err := json.NewDecoder(r.Body).Decode(&purchase); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := connection.Purchases.InsertOne(ctx, purchase)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (connection Connection) DeletePurchaseEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := connection.Purchases.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

// IS VALIDATION NEEDED FOR PUBLIC ADDRESS FORMAT? WE GET IT FROM WEB3.JS ANYWAY
func (connection Connection) FindOrCreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	params := mux.Vars(r)
	publicAddress, _ := params["publicAddress"]
	var user User
	var result bson.M
	filter := bson.D{{"publicAddress", publicAddress}}
	update := bson.D{{"$set", bson.D{{"publicAddress", publicAddress}}}}
	json.NewDecoder(r.Body).Decode(&user)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := connection.Users.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
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
	purchasesCollection := marketplaceDatabase.Collection("purchases") // nest reviews in here
	usersCollection := marketplaceDatabase.Collection("users")

	//document := Listing{
	//	Name:        "Guuuuccciii Gang",
	//	User:        "Nic Raboy",
	//	Description: "Super cool Gucci tshirt with a fucking space eagle",
	//	Price:       150,
	//	Photo:       "photoURL",
	//}

	//listingResult, err := listingsCollection.InsertOne(ctx, document)

	//purchasesResult, err := purchasesCollection.InsertOne(ctx, bson.D{
	//	{"Listing", listingResult.InsertedID},
	//	{"Buyer", "Rich boy"},
	//	{"Seller", "Nic Raboy"},
	//})

	//if err != nil {
	//	log.Fatal(err)
	//}

	// fmt.Printf("%v\n", purchasesResult)

	cur, err := listingsCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	connection := Connection{Listings: listingsCollection, Purchases: purchasesCollection, Users: usersCollection}

	router := mux.NewRouter()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/listing", connection.CreateListingEndpoint).Methods("POST", "OPTIONS")
	router.HandleFunc("/listings", connection.GetListingsEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/listing/{id}", connection.GetListingEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/listing/{id}", connection.UpdateListingEndpoint).Methods("PUT", "OPTIONS")
	router.HandleFunc("/listing/{id}", connection.DeleteListingEndpoint).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/listing/{id}/purchase", connection.CreatePurchaseEndpoint).Methods("POST", "OPTIONS")
	router.HandleFunc("/purchase/{id}", connection.DeletePurchaseEndpoint).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/user/{publicAddress}", connection.FindOrCreateUserEndpoint).Methods("GET", "OPTIONS")
	http.ListenAndServe(":8080", handlers.CORS(header, methods, origins)(router))
}
