package app

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	Id             int     `json:"ID"`
	Name           string  `json:"NAME"`
	Mark           Marks   `json:"MARKS"`
	Total          int     `json:"TOTAL"`
	Average        float32 `json:"AVEARGE"`
	Grade          string  `json:"GRADE"`
	FailCount      int     `json:"FAILED SUBJECT COUNT"`
	FailedSubjects string  `json:"FAILEDSUBJECTS"`
}

type Marks struct {
	Mark1 int `json:"MARK1"`
	Mark2 int `json:"MARK2"`
	Mark3 int `json:"MARK3"`
	Mark4 int `json:"MARK4"`
	Mark5 int `json:"MARk5"`
}

var student []Student

func Get(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(student)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["ID"]
	paramID, _ := strconv.Atoi(params)
	for _, stud := range student {
		if stud.Id == paramID {
			json.NewEncoder(w).Encode(stud)
			return
		}
	}
	json.NewEncoder(w).Encode("No student found with given ID")

}
func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewDecoder(r.Body).Decode(&student)
outer:
	for i := 0; i < len(student); i++ {
		val := reflect.ValueOf(student[i].Mark)
		vals := make([]interface{}, val.NumField())
		for k := 0; k < 5; k++ {
			vals[k] = val.Field(k).Interface()
			//iAreaId := values[i].(int)-->syntax to convert interface type to int
			marks, _ := vals[k].(int)
			if marks > 100 {
				json.NewEncoder(w).Encode("Marks cannot be greater than 100")
				break outer
			} else if marks < 0 {
				json.NewEncoder(w).Encode("Marks cannot be Negative")
				break outer
			}
		}
		student[i].Total = student[i].Mark.Mark1 + student[i].Mark.Mark2 + student[i].Mark.Mark3 + student[i].Mark.Mark4 + student[i].Mark.Mark5
		student[i].Average = float32(student[i].Total) / 5
		if student[i].Average > 90 {
			student[i].Grade = "GRADE O"
		} else if student[i].Average > 80 {
			student[i].Grade = "GRADE A"
		} else if student[i].Average > 70 {
			student[i].Grade = "GRADE B"
		} else if student[i].Average > 60 {
			student[i].Grade = "GRADE C"
		} else if student[i].Average > 50 {
			student[i].Grade = "GRADE D"
		} else {
			student[i].Grade = "GRADE F"
		}
		v := reflect.ValueOf(student[i].Mark)
		values := make([]interface{}, v.NumField())
		for j := 0; j < 5; j++ {
			values[j] = v.Field(j).Interface()
			//iAreaId := values[i].(int)-->syntax to convert interface type to int
			mark, _ := values[j].(int)
			if mark < 50 {
				student[i].FailCount += 1
				switch j {
				case 0:
					student[i].FailedSubjects += "Mark1 "
				case 1:
					student[i].FailedSubjects += "Mark2 "
				case 2:
					student[i].FailedSubjects += "Mark3 "
				case 3:
					student[i].FailedSubjects += "Mark4 "
				case 4:
					student[i].FailedSubjects += "Mark5 "
				}
			}
		}
	}
	json.NewEncoder(w).Encode(student)
}
