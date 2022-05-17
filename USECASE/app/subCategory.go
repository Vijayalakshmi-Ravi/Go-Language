package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SubCategory struct {
	CategoryID      int    `json:"categoryid" bson:"categoryid"`
	SubCategoryID   int    `json:"subcategoryid" bson:"subcategoryid"`
	SubCategoryName string `json:"subcategoryname" bson:"subcategoryname"`
	Description     string `json:"description" bson:"description"`
	Status          string `json:"status" bson:"status"`
	CreatedBy       string `json:"createdby" bson:"createdby"`
	LastUpdatedBy   string `json:"lastupdatedby" bson:"lastupdatedby"`
}
type ResponseSub struct {
	//ErrorMessage  string `json:"error message"`
	StatusCode    int    `json:"statuscode"`
	Status        bool   `json:"status"`
	CustomMessage string `json:"custommmessage"`
}

type ResponseGetSub struct {
	StatusCode    int         `json:"statuscode"`
	Status        bool        `json:"status"`
	CustomMessage string      `json:"custommmessage"`
	Result        SubCategory `json:"Document"`
}

var subCategoryCollection = db().Database("USECASE").Collection("SubCategory")

func convert(event interface{}) SubCategory {
	c := SubCategory{}
	mapstructure.Decode(event, &c)
	fmt.Println(event, c)
	return c
}

func CreateSubCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var subcategory SubCategory
	err := json.NewDecoder(r.Body).Decode(&subcategory)
	if err != nil {
		fmt.Print(err)
	}
	var result primitive.M
	errs := categoryCollection.FindOne(context.TODO(), bson.D{{"categoryid", subcategory.CategoryID}, {"status", "InActive"}}).Decode(&result)
	fmt.Println(errs)
	if errs != nil {
		err1 := subCategoryCollection.FindOne(context.TODO(), bson.D{{"subcategoryid", subcategory.SubCategoryID}}).Decode(&result)
		fmt.Println(err1)
		if err1 != nil {
			insertResult, err := subCategoryCollection.InsertOne(context.TODO(), subcategory)
			if err != nil {
				log.Fatal(err)
			}
			res := ResponseSub{
				StatusCode:    200,
				Status:        true,
				CustomMessage: "Record Inserted Successfully",
			}
			json.NewEncoder(w).Encode(res)
			fmt.Println("Inserted a single document: ", insertResult)
		} else {
			msg := ResponseError{
				ErrorMessage:  "nil",
				StatusCode:    200,
				Status:        false,
				CustomMessage: "SubCategory already Exists",
			}

			json.NewEncoder(w).Encode(msg)
		}
	} else {
		msg := ResponseError{
			ErrorMessage:  "nil",
			StatusCode:    200,
			Status:        false,
			CustomMessage: "Category you selected nolonger available",
		}

		json.NewEncoder(w).Encode(msg)
	}

}

func GetSubCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var body SubCategory
	params := mux.Vars(r)["subcategoryid"]
	id, _ := strconv.Atoi(params)
	err := subCategoryCollection.FindOne(context.TODO(), bson.M{"subcategoryid": id}).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := ResponseError{
			ErrorMessage:  err.Error(),
			StatusCode:    200,
			Status:        false,
			CustomMessage: "Invalid SubCategoryID",
		}

		json.NewEncoder(w).Encode(msg)
		return
	}
	res := ResponseGetSub{
		StatusCode:    200,
		Status:        true,
		CustomMessage: "Success",
		Result:        body,
	}
	json.NewEncoder(w).Encode(res)

}

func UpdateSubCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		SubCategoryID   int    `json:"subcategoryid"`
		CategoryID      int    `json:"categoryid"`
		SubCategoryName string `json:"subcategoryname"`
		Description     string `json:"description"`
		Status          string `json:"status"`
		LastUpdatedBy   string `json:"lastupdatedby"`
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"subcategoryid", body.SubCategoryID}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"categoryid", body.CategoryID}, {"subcategoryname", body.SubCategoryName}, {"description", body.Description}, {"status", body.Status}, {"lastupdatedby", body.LastUpdatedBy}}}}
	var record primitive.M
	errs := categoryCollection.FindOne(context.TODO(), bson.D{{"categoryid", body.CategoryID}, {"status", "InActive"}}).Decode(&record)
	if errs != nil {
		updateResult := subCategoryCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

		var result primitive.M
		_ = updateResult.Decode(&result)
		if result != nil {
			res := Response{
				StatusCode:    200,
				Status:        true,
				CustomMessage: "Updated Successfully",
			}
			json.NewEncoder(w).Encode(res)
		} else {
			msg := ResponseError{
				ErrorMessage:  "nil",
				StatusCode:    200,
				Status:        false,
				CustomMessage: "Update SubCategoryID does not Exists",
			}
			json.NewEncoder(w).Encode(msg)
		}
	} else {
		msg := ResponseError{
			ErrorMessage:  "nil",
			StatusCode:    200,
			Status:        false,
			CustomMessage: "Updating CategoryID no longer available",
		}
		json.NewEncoder(w).Encode(msg)
	}
}

func DeleteSubCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]

	id, _ := strconv.Atoi(params)
	res, err := subCategoryCollection.DeleteOne(context.TODO(), bson.D{{"subcategoryid", id}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount)

}

func GetAllSubCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M
	cur, err := subCategoryCollection.Find(context.TODO(), bson.D{{}})
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
		res := ResponseGetAll{
			StatusCode:    200,
			Status:        true,
			CustomMessage: "Success",
			Result:        results,
		}
		json.NewEncoder(w).Encode(res)
	}
}
