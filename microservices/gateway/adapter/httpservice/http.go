package httpservice

import (
	"log"
	"net/http"

	"github.com/booscaaa/rtdd-golang/microservices/gateway/adapter/httpservice/personservice"
	"github.com/booscaaa/rtdd-golang/microservices/gateway/adapter/httpservice/productservice"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

// Run .
func Run() {
	connPerson, _ := grpc.Dial("person:1111", grpc.WithInsecure())
	connProduct, _ := grpc.Dial("product:2222", grpc.WithInsecure())

	personService := personservice.NewPersonService(connPerson)
	productService := productservice.NewProductService(connProduct)

	router := mux.NewRouter()

	router.Handle("/person", http.HandlerFunc(personService.Fetch)).Methods("GET")
	router.Handle("/person/{id}", http.HandlerFunc(personService.GetByID)).Methods("GET")

	router.Handle("/product", http.HandlerFunc(productService.Fetch)).Methods("GET")

	log.Printf("RUNNING ON PORT 5000")
	http.ListenAndServe(":5000", router)
}
