package person

import (
	"encoding/json"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type deletePersonResponse struct {
	deleted bool
}

func DeletePerson(writer http.ResponseWriter, request *http.Request) {
	var params = mux.Vars(request)
	var idString, _ = uuid.FromString(params["id"])
	var deleted bool = false

	for index, item := range People {
		if item.Id == idString {
			People = append(People[:index], People[index+1:]...)
			deleted = true
		}
	}

	if !deleted {
		json.NewEncoder(writer).Encode("User not found")

		return
	}

	json.NewEncoder(writer).Encode(deletePersonResponse{deleted})
}