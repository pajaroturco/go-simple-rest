package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func HomeEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("Primer api rest")
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Person{})
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = uuid.New().String()
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)

}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: uuid.New().String(), FirstName: "Ryan", LastName: "Park", Address: &Address{
		City: "Dublin", State: "California"}})
	people = append(people, Person{ID: uuid.New().String(), FirstName: "Pierce", LastName: "Brosnan", Address: &Address{
		City: "Reynosa", State: "Tamaulipas"}})

	// endpoints
	router.HandleFunc("/", HomeEndpoint).Methods("GET")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))

}
