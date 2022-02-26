package person_service

import (
	"encoding/json"
	"net/http"
)

func (service service) Fetch(response http.ResponseWriter, request *http.Request) {
	people, err := service.usecase.Fetch()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(response).Encode(people)
}
