package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	// "net/url"
	"os"
	"time"
	"html/template"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

var tpl = template.Must(template.ParseFiles("static/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	listing :=
	tpl.Execute(w, &Listing{})
}

type Listing struct {
	ID primitive.ObjectID	`bson:"_id,omitempty"`
	Name string	`bson:"name,omitempty"`
	User string `bson:"user,omitempty"` // crypto address?
	Description string `bson:"description,omitempty"`
	Price float64 `bson:"price,omitempty"`
	Photo string `bson:"photo,omitempty"` // url?? hosted somewhere? Cloudinary?
}

func (l *Listing) returnFirst() Listing {

}

type Purchase struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Listing Listing `bson:"listing,omitempty"`
	Buyer string `bson:"buyer,omitempty"`
	Seller string `bson:"seller,omitempty"` // crypto address?
}

func goDotEnvVariable(key string) string {

	// load .env file
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


	listingResult, err := listingsCollection.InsertOne(ctx, bson.D{
		{"Name", "Guuuuccciii Gang"},
		{"User", "Nic Raboy"},
		{"Description", "Super cool Gucci tshirt with a fucking space eagle"},
		{"Price", 150},
		{"Photo", "photoURL"},
	})

	fmt.Printf("%v\n", listingResult)

	purchasesResult, err := purchasesCollection.InsertOne(ctx, bson.D{
		{"Listing", listingResult.InsertedID},
		{"Buyer", "Rich boy"},
		{"Seller", "Nic Raboy"},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", purchasesResult)

	// create services


	// attach to HTTP handler
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}