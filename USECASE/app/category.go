package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// struct for storing data
type Category struct {
	CategoryID    int    `json:"categoryid" bson:"categoryid"`
	CategoryName  string `json:"categoryname" bson:"categoryname"`
	Description   string `json:"description" bson:"description"`
	Status        string `json:"status" bson:"status"`
	CreatedBy     string `json:"createdby" bson:"createdby"`
	LastUpdatedBy string `json:"lastupdatedby" bson:"lastupdatedby"`
}

type ResponseError struct {
	ErrorMessage  string `json:"error message"`
	StatusCode    int    `json:"status code"`
	Status        bool   `json:"status"`
	CustomMessage string `json:"customm message"`
}

type Response struct {
	//ErrorMessage  string `json:"error message"`
	StatusCode    int    `json:"status code"`
	Status        bool   `json:"status"`
	CustomMessage string `json:"customm message"`
}

var categoryCollection = db().Database("USECASE").Collection("Category")

func CreateCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		fmt.Print(err)
	}
	var result primitive.M
	err1 := categoryCollection.FindOne(context.TODO(), bson.D{{"categoryid", category.CategoryID}}).Decode(&result)
	if err1 != nil {
		insertResult, err := categoryCollection.InsertOne(context.TODO(), category)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a single document: ", insertResult)
		res := Response{
			StatusCode:    200,
			Status:        true,
			CustomMessage: "Record Inserted Successfully",
		}
		json.NewEncoder(w).Encode(res)
		json.NewEncoder(w).Encode(insertResult.InsertedID)
	} else {
		msg := ResponseError{
			ErrorMessage:  "nil",
			StatusCode:    200,
			Status:        false,
			CustomMessage: "Category already Exists",
		}

		json.NewEncoder(w).Encode(msg)
	}

}

func GetCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var body Category
	params := mux.Vars(r)["categoryid"]
	id, _ := strconv.Atoi(params)
	err := categoryCollection.FindOne(context.TODO(), bson.M{"categoryid": id}).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := ResponseError{
			ErrorMessage:  err.Error(),
			StatusCode:    200,
			Status:        false,
			CustomMessage: "Invalid call",
		}

		json.NewEncoder(w).Encode(msg)
		return
	}
	res := Response{
		StatusCode:    200,
		Status:        true,
		CustomMessage: "Success",
	}
	json.NewEncoder(w).Encode(res)
	json.NewEncoder(w).Encode(body)

}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		CategoryID    int    `json:"categoryid"`
		CategoryName  string `json:"categoryname"`
		Description   string `json:"description"`
		Status        string `json:"status"`
		LastUpdatedBy string `json:"lastupdatedby"`
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"categoryid", body.CategoryID}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"categoryname", body.CategoryName}, {"description", body.Description}, {"status", body.Status}, {"lastupdatedby", body.LastUpdatedBy}}}}
	updateResult := categoryCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)
	json.NewEncoder(w).Encode(result)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["categoryid"]
	id, _ := strconv.Atoi(params)
	res, err := categoryCollection.DeleteOne(context.TODO(), bson.D{{"categoryid", id}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount)

}

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M
	cur, err := categoryCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println(err)
	}
	for cur.Next(context.TODO()) {

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	cur.Close(context.TODO())
	json.NewEncoder(w).Encode(results)
}
