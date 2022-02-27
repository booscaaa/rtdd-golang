package personservice

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (service service) GetByID(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	person, err := service.usecase.GetByID(id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(response).Encode(person)
}
