package person

import (
	"encoding/json"
	"github.com/AlexandreResende/go_agenda_api/src/domain/entities"
	"net/http"
)

var People []entities.Person = make([]entities.Person, 0)

func GetPeople(writer http.ResponseWriter, request *http.Request) {
	if len(People) == 0 {
		json.NewEncoder(writer).Encode(&entities.Person{})

		return
	}

	json.NewEncoder(writer).Encode(People)
}
