package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	// "net/url"
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

// NEED TO ADD JSON

type Listing struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,string,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	User        string             `bson:"user,omitempty" json:"user,omitempty"` // crypto address?
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Price       float64            `bson:"price,omitempty" json:"price, string, omitempty"`
	Photo       string             `bson:"photo,omitempty" json:"photo,omitempty"` // url?? hosted somewhere? Cloudinary?
}

type Purchase struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Listing primitive.ObjectID `bson:"listing,omitempty"`
	Buyer   string             `bson:"buyer,omitempty"`
	Seller  string             `bson:"seller,omitempty"` // crypto address?
}

type Connection struct {
	Listings  *mongo.Collection
	Purchases *mongo.Collection
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
	w.Header().Set("content-type", "application/json")
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

	document := Listing{
		Name:        "Guuuuccciii Gang",
		User:        "Nic Raboy",
		Description: "Super cool Gucci tshirt with a fucking space eagle",
		Price:       150,
		Photo:       "photoURL",
	}

	listingResult, err := listingsCollection.InsertOne(ctx, document)

	purchasesResult, err := purchasesCollection.InsertOne(ctx, bson.D{
		{"Listing", listingResult.InsertedID},
		{"Buyer", "Rich boy"},
		{"Seller", "Nic Raboy"},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", purchasesResult)

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

	connection := Connection{Listings: listingsCollection, Purchases: purchasesCollection}

	router := mux.NewRouter()
	router.HandleFunc("/listing", connection.CreateListingEndpoint).Methods("POST")
	router.HandleFunc("/listings", connection.GetListingsEndpoint).Methods("GET")
	router.HandleFunc("/listing/{id}", connection.UpdateListingEndpoint).Methods("PUT")
	router.HandleFunc("/listing/{id}", connection.DeleteListingEndpoint).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
