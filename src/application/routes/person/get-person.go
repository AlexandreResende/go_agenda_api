package person

import (
	"encoding/json"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

type NotFoundResponse struct {
	message string
}

func GetPerson(writer http.ResponseWriter, request *http.Request) {
	var params = mux.Vars(request)

	log.Println()
	log.Println(params["id"])

	for _, item := range People {
		var idString, _ = uuid.FromString(params["id"])

		log.Println()
		log.Println(idString)

		if item.Id == idString {
			json.NewEncoder(writer).Encode(item)

			return
		}
	}

	json.NewEncoder(writer).Encode(NotFoundResponse{"User not found"})
}