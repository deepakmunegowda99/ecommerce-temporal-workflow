package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pborman/uuid"
)

type Payment struct {
	Status bool   `json:"status"`
	ID     string `json:"id"`
}

func Pay(w http.ResponseWriter, r *http.Request) {

	payment := Payment{
		Status: true,
		ID:     uuid.New(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}

func main() {

	log.Println("Server started on: http://localhost:4003")
	http.HandleFunc("/payment", Pay)
	http.ListenAndServe(":4003", nil)

}
