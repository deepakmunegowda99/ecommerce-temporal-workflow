package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Output struct {
	OriginalPrice   string `json:"original_price"`
	DiscountedPrice string `json:"discounted_price"`
}

func Offer(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("value")

	if value == "" {
		value = "1000"
	}

	min := 5
	max := 12
	discount := rand.Intn(max-min) + min

	i, _ := strconv.Atoi(value)

	final := i - ((i * discount) / 100)

	result := &Output{
		OriginalPrice:   value,
		DiscountedPrice: strconv.Itoa(final),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func main() {

	log.Println("Server started on: http://localhost:4002")
	http.HandleFunc("/offer", Offer)
	http.ListenAndServe(":4002", nil)

}
