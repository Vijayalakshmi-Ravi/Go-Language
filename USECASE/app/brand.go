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

type Brand struct {
	CategoryID    int    `json:"categoryid" bson:"categoryid"`
	SubCategoryID int    `json:"subcategoryid" bson:"subcategoryid"`
	BrandID       int    `json:"brandid" bson:"brandid"`
	BrandName     string `json:"brandname" bson:"brandname"`
	Description   string `json:"description" bson:"description"`
	Status        string `json:"status" bson:"status"`
	CreatedBy     string `json:"createdby" bson:"createdby"`
	LastUpdatedBy string `json:"lastupdatedby" bson:"lastupdatedby"`
}

var brandCollection = db().Database("USECASE").Collection("Brand")

func CreateBrand(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var brand Brand
	err := json.NewDecoder(r.Body).Decode(&brand)
	if err != nil {
		fmt.Print(err)
	}
	var result primitive.M
	errs := categoryCollection.FindOne(context.TODO(), bson.D{{"categoryid", brand.CategoryID}, {"status", "InActive"}}).Decode(&result)
	fmt.Println(errs)
	if errs != nil {
		err1 := subCategoryCollection.FindOne(context.TODO(), bson.D{{"subcategoryid", brand.SubCategoryID}, {"status", "InActive"}}).Decode(&result)
		if err1 != nil {
			err2 := brandCollection.FindOne(context.TODO(), bson.D{{"brandid", brand.BrandID}}).Decode(&result)
			if err2 != nil {
				insertResult, err := brandCollection.InsertOne(context.TODO(), brand)
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
					CustomMessage: "Brand already Exists",
				}

				json.NewEncoder(w).Encode(msg)
			}
		} else {
			msg := ResponseError{
				ErrorMessage:  "nil",
				StatusCode:    200,
				Status:        false,
				CustomMessage: "SubCategory you selected not available",
			}

			json.NewEncoder(w).Encode(msg)
		}
	} else {
		msg := ResponseError{
			ErrorMessage:  "nil",
			StatusCode:    200,
			Status:        false,
			CustomMessage: "Category you selected not available",
		}

		json.NewEncoder(w).Encode(msg)

	}
}

func GetBrand(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var body Brand
	params := mux.Vars(r)["brandid"]
	fmt.Println(params)
	id, _ := strconv.Atoi(params)
	err := brandCollection.FindOne(context.TODO(), bson.M{"brandid": id}).Decode(&body)
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

func UpdateBrand(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		BrandID       int    `json:"brandid"`
		BrandName     string `json:"brandname"`
		Description   string `json:"description"`
		Status        string `json:"status"`
		LastUpdatedBy string `json:"lastupdatedby"`
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"brandid", body.BrandID}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"brandname", body.BrandName}, {"description", body.Description}, {"status", body.Status}, {"lastupdatedby", body.LastUpdatedBy}}}}
	updateResult := brandCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)
	json.NewEncoder(w).Encode(result)
}

func DeleteBrand(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(params)
	res, err := brandCollection.DeleteOne(context.TODO(), bson.D{{"brandid", id}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount)

}

func GetAllBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M
	cur, err := brandCollection.Find(context.TODO(), bson.D{{}})
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
