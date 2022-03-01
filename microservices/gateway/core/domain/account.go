package domain

import (
	"net/http"
)

type AccountService interface {
	Login(response http.ResponseWriter, request *http.Request)
}
