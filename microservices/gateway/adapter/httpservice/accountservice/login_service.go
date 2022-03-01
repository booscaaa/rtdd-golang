package accountservice

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/booscaaa/rtdd-golang/microservices/gateway/core/domain"
)

func (service service) Login(response http.ResponseWriter, request *http.Request) {
	grpcReponse, err := service.grpc.Login(context.Background(), &domain.LoginAccountRequest{
		User:     "admin",
		Password: "admin",
	})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(grpcReponse)
}
