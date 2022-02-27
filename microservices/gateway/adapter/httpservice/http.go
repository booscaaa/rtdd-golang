package httpservice

import (
	"log"
	"net/http"

	"github.com/booscaaa/rtdd-golang/microservices/gateway/adapter/httpservice/personservice"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

// Run .
func Run() {
	conn, _ := grpc.Dial(
		"localhost:1111",
		grpc.WithInsecure(),
	)

	personService := personservice.NewPersonService(conn)

	router := mux.NewRouter()

	router.Handle("/person", http.HandlerFunc(personService.Fetch)).Methods("GET")
	router.Handle("/person/{id}", http.HandlerFunc(personService.GetByID)).Methods("GET")

	log.Printf("RUNNING ON PORT 5000")
	http.ListenAndServe(":5000", router)
}
