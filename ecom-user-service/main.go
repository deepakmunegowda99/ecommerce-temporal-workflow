package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pallinder/go-randomdata"
)

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func UserDetails(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	data := User{
		ID:      id,
		Name:    randomdata.FirstName(randomdata.Male),
		Email:   randomdata.Email(),
		Phone:   randomdata.PhoneNumber(),
		Address: randomdata.Address(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func main() {

	log.Println("Server started on: http://localhost:4000")
	http.HandleFunc("/user", UserDetails)
	http.ListenAndServe(":4000", nil)

}
