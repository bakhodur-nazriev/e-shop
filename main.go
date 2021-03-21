package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Product struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Size        string  `json:"size"`
	Brand       string  `json:"brand"`
}

var Products []Product

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func singleProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, product := range Products {
		if product.Id == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}
func allProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All products Endpoint")
	json.NewEncoder(w).Encode(Products)

}
func createProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var product Product
	json.Unmarshal(reqBody, &product)

	Products = append(Products, product)
	json.NewEncoder(w).Encode(product)
}
func updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	//reqBody, _ := ioutil.ReadAll(r.Body)

	for i, product := range Products {
		if product.Id == id {
			Products = append(Products[:i], Products[i+1:]...)
			json.NewEncoder(w).Encode(product)
		}
	}
}
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, product := range Products {
		if product.Id == id {
			Products = append(Products[:i], Products[i+1:]...)
			json.NewEncoder(w).Encode(product)
		}
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/products", allProducts).Methods("GET")
	router.HandleFunc("/product/{id}", singleProduct).Methods("GET")
	router.HandleFunc("/product", createProduct).Methods("POST")
	router.HandleFunc("/product/{id}", updateProduct).Methods("PUT")
	router.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {

	Products = []Product{
		{
			Id:          "1",
			Title:       "Shoes",
			Description: "One of the best",
			Price:       50.5,
			Size:        "XXL",
			Brand:       "Nike",
		},
		{
			Id:          "2",
			Title:       "Shirt",
			Description: "The coolest shirt",
			Price:       36.2,
			Size:        "XL",
			Brand:       "ESSE",
		},
	}

	handleRequests()
}
