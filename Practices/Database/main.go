package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	usersCollection := client.Database("testing1").Collection("users")
	//dropps all the existing documents in the collections and makes empty
	if err = usersCollection.Drop(context.TODO()); err != nil {
		panic(err)
	}
	// insert a single document into a collection
	// create a bson.D object
	user := bson.D{primitive.E{Key: "fullName", Value: "User 1"}, {Key: "age", Value: 25}}
	// insert the bson object using InsertOne()
	result, err := usersCollection.InsertOne(context.TODO(), user)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the id of the newly inserted object
	fmt.Println(result.InsertedID)

	// insert multiple documents into a collection
	// create a slice of bson.D objects
	users := []interface{}{
		bson.D{primitive.E{Key: "fullName", Value: "User 2"}, {Key: "age", Value: 25}},
		bson.D{primitive.E{Key: "fullName", Value: "User 3"}, {Key: "age", Value: 20}},
		bson.D{primitive.E{Key: "fullName", Value: "User 4"}, {Key: "age", Value: 28}},
	}
	// insert the bson object slice using InsertMany()
	results, err := usersCollection.InsertMany(context.TODO(), users)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the ids of the newly inserted objects
	fmt.Println(results.InsertedIDs)

	// retrieve all the documents in a collection
	cursor, err := usersCollection.Find(context.TODO(), bson.D{})
	// check for errors in the finding
	if err != nil {
		panic(err)
	}

	// convert the cursor result to bson
	var result1 []bson.M
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &result1); err != nil {
		panic(err)
	}

	// display the documents retrieved
	fmt.Println("displaying all results in a collection")
	for _, result := range result1 {
		fmt.Println(result)
	}

	//-----------------------------------------from here filtering-------------------------------------------------

	// retrieve single and multiple documents with a specified filter using FindOne() and Find()
	// create a search filer
	filter := bson.D{primitive.E{
		Key: "$and",
		Value: bson.A{
			bson.D{
				primitive.E{Key: "age", Value: bson.D{primitive.E{Key: "$gt", Value: 20}}},
			},
		},
	},
	}

	// retrieve all the documents that match the filter
	cursor1, err1 := usersCollection.Find(context.TODO(), filter)
	// check for errors in the finding
	if err1 != nil {
		panic(err1)
	}

	// convert the cursor result to bson
	var result2 []bson.M
	// check for errors in the conversion
	if err = cursor1.All(context.TODO(), &result2); err1 != nil {
		panic(err1)
	}

	// display the documents retrieved
	fmt.Println("displaying all results from the search query")
	for _, result := range result2 {
		fmt.Println(result)
	}

	// retrieving the first document that match the filter
	var result3 bson.M
	// check for errors in the finding
	if err1 = usersCollection.FindOne(context.TODO(), filter).Decode(&result3); err1 != nil {
		panic(err1)
	}

	// display the document retrieved
	fmt.Println("displaying the first result from the search filter")
	fmt.Println(result3)

	//----------------------------------Updation ------------------------------------------
	user5 := bson.D{primitive.E{Key: "fullName", Value: "User 5"}, {Key: "age", Value: 22}}
	//insert one document
	insertResult, err := usersCollection.InsertOne(context.TODO(), user5)

	if err != nil {
		panic(err)
	}

	//create the update query for the client
	update := bson.D{primitive.E{
		Key: "$set",
		Value: bson.D{
			primitive.E{Key: "fullName", Value: "User V"}}},
		primitive.E{Key: "$inc", Value: bson.D{primitive.E{
			Key:   "age",
			Value: 1}}},
	}

	// execute the UpdateByID() function with the filter and update query
	resultu, erru := usersCollection.UpdateByID(context.TODO(), insertResult.InsertedID, update)
	// check for errors in the updating
	if erru != nil {
		panic(erru)
	}
	// display the number of documents updated
	fmt.Println("Number of documents updated:", resultu.ModifiedCount)
	//-------------------------------------------------------------------------------------------
	//updation of 4 document updating full name from user 4 =>user four
	//document 4 updation
	updateU4 := bson.D{primitive.E{
		Key: "$set",
		Value: bson.D{
			primitive.E{Key: "fullName", Value: "User three"}}},
		primitive.E{Key: "$inc", Value: bson.D{primitive.E{
			Key:   "age",
			Value: 1}}},
	}

	resultup, errup := usersCollection.UpdateOne(context.TODO(), bson.D{primitive.E{Key: "fullName", Value: "User 3"}}, updateU4)
	if errup != nil {
		panic(errup)
	}
	fmt.Println("Records Updated: ", resultup.ModifiedCount)
}
