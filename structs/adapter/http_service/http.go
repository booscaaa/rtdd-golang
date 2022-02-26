package http_service

import (
	"log"
	"net/http"

	"github.com/booscaaa/rtdd-golang/structs/adapter/sqlite"
	"github.com/booscaaa/rtdd-golang/structs/di"
	"github.com/gorilla/mux"
)

func Run() {
	db := sqlite.GetConnection()
	sqlite.ConfigConnection(db)

	personService := di.ConfigPersonInjection(db)

	router := mux.NewRouter()

	router.Handle("/person", http.HandlerFunc(personService.Fetch)).Methods("GET")
	router.Handle("/person/{id}", http.HandlerFunc(personService.GetByID)).Methods("GET")

	log.Printf("RUNNING ON PORT 4000")
	http.ListenAndServe(":4000", router)
}
