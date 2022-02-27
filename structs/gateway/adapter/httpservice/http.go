package httpservice

import (
	"log"
	"net/http"

	"github.com/booscaaa/rtdd-golang/gateway/adapter/sqlite"
	"github.com/booscaaa/rtdd-golang/gateway/di"
	"github.com/gorilla/mux"
)

// Run .
func Run() {
	db := sqlite.GetConnection()
	sqlite.ConfigConnection(db)

	personService := di.ConfigPersonServiceInjection(db)

	router := mux.NewRouter()

	router.Handle("/person", http.HandlerFunc(personService.Fetch)).Methods("GET")
	router.Handle("/person/{id}", http.HandlerFunc(personService.GetByID)).Methods("GET")

	log.Printf("RUNNING ON PORT 5000")
	http.ListenAndServe(":5000", router)
}
