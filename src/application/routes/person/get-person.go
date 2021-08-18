package person

import (
	"encoding/json"
	"github.com/AlexandreResende/go_agenda_api/src/domain/entities"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func GetPerson(writer http.ResponseWriter, request *http.Request) {
	var params = mux.Vars(request)

	for _, item := range People {
		var idString, _ = uuid.FromString(params["id"])

		if item.Id == idString {
			json.NewEncoder(writer).Encode(item)

			return
		}
	}

	json.NewEncoder(writer).Encode(&entities.Person{})
}