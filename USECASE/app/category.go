package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var validate *validator.Validate

// struct for storing data
type Category struct {
	CategoryID    int    `json:"categoryid" bson:"categoryid" validate:"number,requried"`
	CategoryName  string `json:"categoryname" bson:"categoryname" validate:"alpha,requried"`
	Description   string `json:"description" bson:"description" validate:"alpha,requried"`
	Status        string `json:"status" bson:"status" validate:"alpha,requried"`
	CreatedBy     string `json:"createdby" bson:"createdby" validate:"alpha,requried"`
	LastUpdatedBy string `json:"lastupdatedby" bson:"lastupdatedby" validate:"alpha,requried"`
}

type ResponseError struct {
	ErrorMessage  string `json:"errormessage"`
	StatusCode    int    `json:"statuscode"`
	Status        bool   `json:"status"`
	CustomMessage string `json:"custommmessage"`
}
type ResponseValidateError struct {
	ErrorMessage  string              `json:"errormessage"`
	StatusCode    int                 `json:"statuscode"`
	Status        bool                `json:"status"`
	CustomMessage map[string][]string `json:"custommmessage"`
}
type Response struct {
	StatusCode    int    `json:"statuscode"`
	Status        bool   `json:"status"`
	CustomMessage string `json:"custommmessage"`
}
type ResponseGetAll struct {
	//ErrorMessage  string `json:"error message"`
	StatusCode    int           `json:"statuscode"`
	Status        bool          `json:"status"`
	CustomMessage string        `json:"custommmessage"`
	Result        []primitive.M `json:"Documents"`
}

type ResponseGet struct {
	StatusCode    int      `json:"statuscode"`
	Status        bool     `json:"status"`
	CustomMessage string   `json:"custommmessage"`
	Result        Category `json:"Document"`
}

var categoryCollection = db().Database("USECASE").Collection("Category")

func CreateCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		fmt.Print(err)
	}
	validate = validator.New()
	errvalue := validate.Struct(category)
	if errvalue != nil {
		errors := make(map[string][]string)
		for _, err := range errvalue.(validator.ValidationErrors) {
			var name string
			name = strings.ToLower(err.StructField())
			fmt.Println("value", err.Type())
			switch err.Tag() {
			case "required":
				errors[name] = append(errors[name], "The "+name+" is required")
				break
			case "alpha":
				errors[name] = append(errors[name], "The "+name+" should be alpha")
				break
			case "number":
				errors[name] = append(errors[name], "The "+name+" should be equal to the "+err.Param())
				break
			default:
				errors[name] = append(errors[name], "The "+name+" is invalid")
				break
			}
		}
		msg := ResponseValidateError{
			ErrorMessage:  "nill",
			StatusCode:    200,
			Status:        false,
			CustomMessage: errors,
		}
		json.NewEncoder(w).Encode(msg)
		return
	}
	var result primitive.M
	err1 := categoryCollection.FindOne(context.TODO(), bson.D{{"categoryid", category.CategoryID}}).Decode(&result)
	if err1 != nil {
		if category.CategoryID == 0 {
			msg := ResponseError{
				ErrorMessage:  "nil",
				StatusCode:    200,
				Status:        false,
				CustomMessage: "Category ID is not given",
			}
			json.NewEncoder(w).Encode(msg)
		} else {
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
			//json.NewEncoder(w).Encode(insertResult.InsertedID)
		}
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
			CustomMessage: "Invalid CategoryID",
		}

		json.NewEncoder(w).Encode(msg)
		return
	}
	res := ResponseGet{
		StatusCode:    200,
		Status:        true,
		CustomMessage: "Success",
		Result:        body,
	}
	json.NewEncoder(w).Encode(res)

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
			CustomMessage: "Update CategoryID does not Exists",
		}
		json.NewEncoder(w).Encode(msg)
	}

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
