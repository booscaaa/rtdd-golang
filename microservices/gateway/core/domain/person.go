package domain

import (
	"net/http"
)

type PersonService interface {
	Fetch(response http.ResponseWriter, request *http.Request)
	GetByID(response http.ResponseWriter, request *http.Request)
}
