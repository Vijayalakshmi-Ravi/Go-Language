package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	ID   int    `json:"EID"`
	Name string `json:"ENAME"`
	Dept string `json:"DEPT"`
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello good day....")
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	EmployeeDetails := []Employee{
		{ID: 1, Name: "Assheer", Dept: "logistics"},
		{ID: 2, Name: "Crepe", Dept: "Admin"},
		{ID: 3, Name: "Fazil", Dept: "IT"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(EmployeeDetails)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(EmployeeDetails)
	}
}

func getEmpID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["ID"])
}
