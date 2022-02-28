package domain

import (
	"net/http"
)

type ProductService interface {
	Fetch(response http.ResponseWriter, request *http.Request)
}
