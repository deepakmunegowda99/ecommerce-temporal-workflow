package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/wawandco/fako"
)

type Product struct {
	ID          string `json:"id"`
	Company     string `json:"company" fako:"company"`
	Product     string `json:"product" fako:"product"`
	ProductName string `json:"product_name" fako:"product_name"`
	Available   bool   `json:"available"`
	Cost        int    `json:"cost"`
}

func ProductDetails(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	min := 500
	max := 2000

	var product Product
	product.ID = id
	product.Cost = rand.Intn(max-min) + min
	product.Available = true
	fako.Fill(&product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func main() {

	log.Println("Server started on: http://localhost:4001")
	http.HandleFunc("/product", ProductDetails)
	http.ListenAndServe(":4001", nil)

}
