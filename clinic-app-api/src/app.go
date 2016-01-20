package main

import (
	"fmt"
	"clinic"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/api/v1/pets", clinic.GetAllPets).Methods("GET")
	rtr.HandleFunc("/api/v1/pets/{id}", clinic.GetOnePet).Methods("GET")
	rtr.HandleFunc("/api/v1/pets", clinic.CreateNewPet).Methods("POST")
	http.Handle("/", rtr)

	http.ListenAndServe(":8080", nil)
}
