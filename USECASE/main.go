package main

import (
	"log"
	"net/http"
	"usecase/app"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path

	//Routes
	//Category
	s.HandleFunc("/createCategory", app.CreateCategory).Methods("POST")
	s.HandleFunc("/getAllCategory", app.GetAllCategory).Methods("GET")
	s.HandleFunc("/getCategory/{categoryid}", app.GetCategory).Methods("GET")
	s.HandleFunc("/updateCategory", app.UpdateCategory).Methods("PUT")
	s.HandleFunc("/deleteCategory/{categoryid}", app.DeleteCategory).Methods("DELETE")

	//SubCategory
	s.HandleFunc("/createSubCategory", app.CreateSubCategory).Methods("POST")
	s.HandleFunc("/getAllSubCategory", app.GetAllSubCategory).Methods("GET")
	s.HandleFunc("/getSubCategory/{subcategoryid}", app.GetSubCategory).Methods("GET")
	s.HandleFunc("/updateSubCategory", app.UpdateSubCategory).Methods("PUT")
	s.HandleFunc("/deleteSubCategory/{id}", app.DeleteSubCategory).Methods("DELETE")

	//Brand
	s.HandleFunc("/createBrand", app.CreateBrand).Methods("POST")
	s.HandleFunc("/getAllBrand", app.GetAllBrand).Methods("GET")
	s.HandleFunc("/getBrand/{brandid}", app.GetBrand).Methods("GET")
	s.HandleFunc("/updateBrand", app.UpdateBrand).Methods("PUT")
	s.HandleFunc("/deleteBrand/{id}", app.DeleteBrand).Methods("DELETE")

	//Variant
	s.HandleFunc("/createVariant", app.CreateVariant).Methods("POST")
	s.HandleFunc("/getAllVariant", app.GetAllVariant).Methods("GET")
	s.HandleFunc("/getVariant/{variantid}", app.GetVariant).Methods("GET")
	s.HandleFunc("/updateVariant", app.UpdateVariant).Methods("PUT")
	s.HandleFunc("/deleteVariant/{id}", app.DeleteVariant).Methods("DELETE")

	//Products
	s.HandleFunc("/createProduct", app.CreateProduct).Methods("POST")
	s.HandleFunc("/getAllProduct", app.GetAllProduct).Methods("GET")
	s.HandleFunc("/getProduct/{productid}", app.GetProduct).Methods("GET")
	s.HandleFunc("/updateProduct", app.UpdateProduct).Methods("PUT")
	s.HandleFunc("/deleteProduct/{id}", app.DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}
