package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type product struct {
	code  string  `json:"code"`
	name  string  `json:"name"`
	price float64 `json:"price"`
}

type basket []product

func getBasket(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My empty basket!")
}

func createBasket(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/basket", getBasket).Methods("GET")

}
