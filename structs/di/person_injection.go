package di

import (
	"database/sql"

	"github.com/booscaaa/rtdd-golang/structs/adapter/http_service/person_service"
	"github.com/booscaaa/rtdd-golang/structs/adapter/sqlite/person_repository"
	"github.com/booscaaa/rtdd-golang/structs/core/domain"
	"github.com/booscaaa/rtdd-golang/structs/core/person_usecase"
)

func ConfigPersonInjection(db *sql.DB) domain.PersonService {
	personRepository := person_repository.NewPersonRepository(db)
	personUsecase := person_usecase.NewPersonUseCase(personRepository)
	personService := person_service.NewPersonService(personUsecase)

	return personService
}
