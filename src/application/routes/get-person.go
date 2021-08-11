package routes

import (
	"encoding/json"
	"github.com/AlexandreResende/go_agenda_api/src/domain/entities"
	"github.com/gorilla/mux"
	"net/http"
)

func GetPerson(writer http.ResponseWriter, request *http.Request) {
	var params = mux.Vars(request)

	for _, item := range People {
		if item.Id == params["id"] {
			json.NewEncoder(writer).Encode(item)

			return
		}
	}

	json.NewEncoder(writer).Encode(&entities.Person{})
}