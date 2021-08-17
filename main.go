package main

import (
	"github.com/AlexandreResende/go_agenda_api/src/application/routes"
	"github.com/AlexandreResende/go_agenda_api/src/domain/entities"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func main() {
	var router *mux.Router = mux.NewRouter()

	routes.People = append(
		routes.People,
		entities.Person{Id: uuid.NewV4(), FirstName: "John", LastName: "Doe", Address: &entities.Address{City: "City X", State: "State X"}},
	)

	router.HandleFunc("/contato", routes.GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", routes.GetPerson).Methods("GET")
	router.HandleFunc("/contato", routes.CreatePerson).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}