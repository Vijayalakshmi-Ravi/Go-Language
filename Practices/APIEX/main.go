package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Num struct {
	A int `json:"NUM1" `
	B int `json:"NUM2" `
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	n := Num{}
	json.NewDecoder(r.Body).Decode(&n)
	fmt.Println(n)
	sum := n.A + n.B
	json.NewEncoder(w).Encode(sum)

}

func main() {
	fmt.Println("Welcome")
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/test", get).Methods(http.MethodGet)
	api.HandleFunc("/test", post).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", api))

}
