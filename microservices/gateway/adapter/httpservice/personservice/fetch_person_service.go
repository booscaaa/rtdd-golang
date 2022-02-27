package personservice

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/booscaaa/rtdd-golang/microservices/person/core/domain"
)

func (service service) Fetch(response http.ResponseWriter, request *http.Request) {
	grpcReponse, err := service.grpc.Fetch(context.Background(), &domain.FetchRequest{})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(grpcReponse.People)
}
