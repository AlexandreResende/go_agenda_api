package person

import (
	"encoding/json"
	"github.com/AlexandreResende/go_agenda_api/src/domain/entities"
	"github.com/satori/go.uuid"
	"net/http"
)

func CreatePerson(writer http.ResponseWriter, request *http.Request) {
	var person entities.Person
	var people []entities.Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	person.Id = uuid.NewV4()
	people = append(people, person)

	json.NewEncoder(writer).Encode(people)
}