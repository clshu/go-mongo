package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/clshu/go-mongo/api"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty bson:"_id,omitempty"`
	FirstName string             `json:"firstName,omitempty bson:"firstName,omitempty"`
	LastName  string             `json:"lastName,omitempty bson:"lastName,omitempty"`
}

var client *mongo.Client

func main() {
	fmt.Println("Starting the application")
	uri := os.Getenv("MONGODB_URI_GO")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected and pinged.")
	}

	router := mux.NewRouter()
	router.HandleFunc("/app/person/create", api.CreatePerson).Methods("GET")
	http.ListenAndServe(":3030", router)

}
