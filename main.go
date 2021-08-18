package main

import (
	"github.com/AlexandreResende/go_agenda_api/src/application/routes/person"
	"github.com/AlexandreResende/go_agenda_api/src/domain/entities"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func main() {
	var router *mux.Router = mux.NewRouter()

	person.People = append(
		person.People,
		entities.Person{Id: uuid.NewV4(), FirstName: "John", LastName: "Doe", Address: &entities.Address{City: "City X", State: "State X"}},
	)

	router.HandleFunc("/contato", person.GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", person.GetPerson).Methods("GET")
	router.HandleFunc("/contato", person.CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", person.DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}