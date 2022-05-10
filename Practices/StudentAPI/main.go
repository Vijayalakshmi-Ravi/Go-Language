package main

import (
	"STUDENTAPI/app"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	//api.HandleFunc("/test", get).Methods(http.MethodGet)
	api.HandleFunc("/Student", app.Post).Methods(http.MethodPost)
	api.HandleFunc("/Student", app.Get).Methods(http.MethodGet)
	api.HandleFunc("/Student/{ID}", app.GetByID).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", api))

}
