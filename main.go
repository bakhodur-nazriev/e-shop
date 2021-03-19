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

/*Pages handlers function*/
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage!")
}
func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to About Page!")
}
func contactPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Contact Page!")
}

/* Products handler function */
func allProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Products)
}
func product(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, product := range Products {
		if product.Id == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}
func createNewProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var pro Product
	json.Unmarshal(reqBody, &pro)

	Products = append(Products, pro)
	json.NewEncoder(w).Encode(pro)
}
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, p := range Products {
		if p.Id == id {
			Products = append(Products[:i], Products[i:]...)
		}
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	/* Pages routes */
	router.HandleFunc("/", homePage)
	router.HandleFunc("/about", aboutPage)
	router.HandleFunc("/contact", contactPage)

	/* Products routes */
	router.HandleFunc("/products", allProducts)
	router.HandleFunc("/product/{id}", product)
	router.HandleFunc("/product", createNewProduct).Methods("POST")
	router.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {

	Products = []Product{
		{
			Id:          "1",
			Title:       "T-shirt",
			Description: "This is the awesome t-shirt",
			Price:       10.12,
			Size:        "XXL",
			Brand:       "Nike",
		},
		{
			Id:          "2",
			Title:       "Shoes",
			Description: "One of the coolest shoes",
			Price:       250.12,
			Size:        "41",
			Brand:       "ESSE",
		},
	}

	handleRequests()
}
