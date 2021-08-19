package person

import (
	"encoding/json"
	"github.com/AlexandreResende/go_agenda_api/src/domain/entities"
	"net/http"
)

var People []entities.Person = make([]entities.Person, 0)

type GetPeopleResponse struct {
	People []entities.Person `json:"people"`
}

func GetPeople(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(GetPeopleResponse{People})
}
// check json.Unmarshal