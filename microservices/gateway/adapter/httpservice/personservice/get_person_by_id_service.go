package personservice

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/booscaaa/rtdd-golang/microservices/person/core/domain"
	"github.com/gorilla/mux"
)

func (service service) GetByID(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	grpcReponse, err := service.grpc.GetByID(
		context.Background(),
		&domain.GetByIDRequest{Id: int32(id)},
	)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(grpcReponse.Person)
}
