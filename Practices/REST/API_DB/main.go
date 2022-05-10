package main

import (
	"API_DB/app"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		fmt.Println("not connected")
		panic(err)

	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Println("not connected")
		panic(err)
	}
	app.Start()
	// usersCollection := client.Database("Employee").Collection("Employee")
	// //dropps all the existing documents in the collections and makes empty
	// if err = usersCollection.Drop(context.TODO()); err != nil {
	// 	panic(err)
	// }
}
