package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	//gorilla mux
	router := mux.NewRouter()

	//route defining
	router.HandleFunc("/greet", Greet).Methods(http.MethodGet)
	router.HandleFunc("/getAllEmployees", GetAllEmployees).Methods(http.MethodGet)
	router.HandleFunc("/getEmployeeID/{ID:[0-9]+}", getEmpID).Methods(http.MethodGet)

	//starting the srver
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
