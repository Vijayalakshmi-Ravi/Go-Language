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

type Variant struct {
	VariantID     int    `json:"variantid" bson:"variantid"`
	VariantName   string `json:"variantname" bson:"variantname"`
	Description   string `json:"description" bson:"description"`
	Status        string `json:"status" bson:"status"`
	CreatedBy     string `json:"createdby" bson:"createdby"`
	LastUpdatedBy string `json:"lastupdatedby" bson:"lastupdatedby"`
}

var variantCollection = db().Database("USECASE").Collection("Variant")

func CreateVariant(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var variant Variant
	err := json.NewDecoder(r.Body).Decode(&variant)
	if err != nil {
		fmt.Print(err)
	}
	var result primitive.M
	err1 := variantCollection.FindOne(context.TODO(), bson.D{{"variantid", variant.VariantID}}).Decode(&result)
	if err1 != nil {
		insertResult, err := variantCollection.InsertOne(context.TODO(), variant)
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
			CustomMessage: "Variant already Exists",
		}

		json.NewEncoder(w).Encode(msg)
	}
}

func GetVariant(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var body Variant
	params := mux.Vars(r)["variantid"]
	fmt.Println(params)
	id, _ := strconv.Atoi(params)
	err := variantCollection.FindOne(context.TODO(), bson.M{"variantid": id}).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := ResponseError{
			ErrorMessage:  err.Error(),
			StatusCode:    200,
			Status:        false,
			CustomMessage: "Invalid VariantID",
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

func UpdateVariant(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		VariantID     int    `json:"variantid"`
		VariantName   string `json:"variantname"`
		Description   string `json:"description"`
		Status        string `json:"status"`
		LastUpdatedBy string `json:"lastupdatedby"`
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"variantname", body.VariantName}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"variantname", body.VariantName}, {"description", body.Description}, {"status", body.Status}, {"lastupdatedby", body.LastUpdatedBy}}}}
	updateResult := variantCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)
	json.NewEncoder(w).Encode(result)
}

func DeleteVariant(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		fmt.Printf(err.Error())
	}

	res, err := variantCollection.DeleteOne(context.TODO(), bson.D{{"variantid", id}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount)

}

func GetAllVariant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M
	cur, err := variantCollection.Find(context.TODO(), bson.D{{}})
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
	if results == nil {
		msg := ResponseError{
			ErrorMessage:  "nill",
			StatusCode:    200,
			Status:        false,
			CustomMessage: "Empty Collection",
		}
		json.NewEncoder(w).Encode(msg)
	} else {
		cur.Close(context.TODO())
		res := Response{
			StatusCode:    200,
			Status:        true,
			CustomMessage: "Success",
		}
		json.NewEncoder(w).Encode(res)
		json.NewEncoder(w).Encode(results)
	}
}
