package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// struct for storing data
type Employee struct {
	EmpID   int    `json:"EMP_ID" bson:"EMP_ID"`
	EmpName string `json:"EMP_NAME" bson:"EMP_NAME"`
	Dept    string `json:"EMP_DEPT" bson:"EMP_DEPT"`
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {

}

var client *mongo.Client

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("something is wrong in post")
	w.Header().Add("content-type", "application/json")
	var employee Employee
	json.NewDecoder(r.Body).Decode(&employee)
	fmt.Println("something is wrong in post2")
	collection := client.Database("Employee").Collection("Employee")
	fmt.Println("something is wrong in post3")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	fmt.Println("something is wrong in post4")
	result, _ := collection.InsertOne(ctx, employee)
	json.NewEncoder(w).Encode(&result)
}
