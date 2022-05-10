package main

import (
	"DBAPIEg/app"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path

	//Routes

	s.HandleFunc("/createProfile", app.CreateProfile).Methods("POST")
	s.HandleFunc("/getAllUsers", app.GetAllUsers).Methods("GET")
	s.HandleFunc("/getUserProfile", app.GetUserProfile).Methods("GET")
	s.HandleFunc("/updateProfile", app.UpdateProfile).Methods("PUT")
	s.HandleFunc("/deleteProfile/{id}", app.DeleteProfile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}
