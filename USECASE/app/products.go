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
type Product struct {
	ProductID   int    `json:"productid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	//Arr           [3]int `json:"Details"`
	Details       map[string]int `json:"Details"`
	VariantID     int            `json:"variantid"`
	Status        string         `json:"status"`
	CreatedBy     string         `json:"createdby"`
	LastUpdatedBy string         `json:"lastupdatedby"`
}

var productCollection = db().Database("USECASE").Collection("Products")

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product Product

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Fatal(err)
	}
	var result primitive.M
	err1 := productCollection.FindOne(context.TODO(), bson.D{{"productid", product.ProductID}}).Decode(&result)
	if err1 != nil {
		insertResult, err := productCollection.InsertOne(context.TODO(), product)
		if err != nil {
			log.Fatal(err)
		}
		res := Response{
			StatusCode:    200,
			Status:        true,
			CustomMessage: "Record Inserted Successfully",
		}
		json.NewEncoder(w).Encode(res)
		fmt.Println("Inserted a single document: ", insertResult)
		json.NewEncoder(w).Encode(insertResult.InsertedID)
	} else {
		msg := ResponseError{
			ErrorMessage:  "nil",
			StatusCode:    200,
			Status:        false,
			CustomMessage: "Product already Exists",
		}

		json.NewEncoder(w).Encode(msg)

	}
}

func GetProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var body Product
	params := mux.Vars(r)["productid"]
	fmt.Println(params)
	id, _ := strconv.Atoi(params)
	err := productCollection.FindOne(context.TODO(), bson.M{"productid": id}).Decode(&body)
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

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		ProductID     int    `json:"productid"`
		Title         string `json:"Title"`
		Description   string `json:"description"`
		Status        string `json:"status"`
		LastUpdatedBy string `json:"lastupdatedby"`
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"productid", body.ProductID}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"title", body.Title}, {"description", body.Description}, {"status", body.Status}, {"lastupdatedby", body.LastUpdatedBy}}}}
	updateResult := productCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)
	json.NewEncoder(w).Encode(result)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		fmt.Printf(err.Error())
	}
	res, err := productCollection.DeleteOne(context.TODO(), bson.D{{"productid", id}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount)

}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M
	cur, err := productCollection.Find(context.TODO(), bson.D{{}})
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
