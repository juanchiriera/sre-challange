package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/mux"
)

type product struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type productList []product

var basket productList

var products = []product{
	{
		Code:  "PEN",
		Name:  "Lana Pen",
		Price: 5.0,
	},
	{
		Code:  "TSHIRT",
		Name:  "Lana T-Shirt",
		Price: 20.0,
	},
	{
		Code:  "MUG",
		Name:  "Lana Coffee Mug",
		Price: 7.50,
	},
}

//API Call to get basket items
func getBasket(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(basket)
}

//API Call to create new basket
func createBasket(w http.ResponseWriter, r *http.Request) {
	basket = productList{}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(basket)
}

//API Call to add product to basket
func addProduct(w http.ResponseWriter, r *http.Request) {
	productCode := mux.Vars(r)["productCode"]
	for _, currentProduct := range products {
		if currentProduct.Code == productCode {
			basket = append(basket, currentProduct)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(basket)
			log.Print(basket)
		}
	}
}

//API Call to get the total value
func getTotal(w http.ResponseWriter, r *http.Request) {
	total := 0.0
	calculateTotal(basket, &total)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Total: %v", total)
}

//Calculate the total value of the items in the basket
func calculateTotal(basket productList, total *float64) {
	penCounter := 0
	tshirtCounter := 0.0
	for _, product := range basket {
		*total += product.Price
		if product.Code == "PEN" {
			penCounter++
		} else if product.Code == "TSHIRT" {
			tshirtCounter++
		}
	}
	if penCounter > 1 {
		*total = *total - (math.Round(float64(penCounter)/2-0.3))*5.0
	}
	if tshirtCounter >= 3 {
		*total = *total - tshirtCounter*5.0
	}
}

//Main application
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/basket", getBasket).Methods("GET")
	router.HandleFunc("/total", getTotal).Methods("GET")
	router.HandleFunc("/basket", createBasket).Methods("POST")
	router.HandleFunc("/product/{productCode}", addProduct).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}
