package di

import (
	"database/sql"

	"github.com/booscaaa/rtdd-golang/microservices/gateway/adapter/cmdservice/personcmd"
	"github.com/booscaaa/rtdd-golang/microservices/gateway/adapter/httpservice/personservice"
	"github.com/booscaaa/rtdd-golang/microservices/gateway/adapter/sqlite/personrepository"
	"github.com/booscaaa/rtdd-golang/microservices/gateway/core/domain"
	"github.com/booscaaa/rtdd-golang/microservices/gateway/core/personusecase"
)

func getPersonUseCase(db *sql.DB) domain.PersonUseCase {
	personRepository := personrepository.NewPersonRepository(db)
	personUsecase := personusecase.NewPersonUseCase(personRepository)

	return personUsecase
}

// ConfigPersonServiceInjection .
func ConfigPersonServiceInjection(db *sql.DB) domain.PersonService {
	personUseCase := getPersonUseCase(db)
	personService := personservice.NewPersonService(personUseCase)

	return personService
}

// ConfigPersonCMDInjection .
func ConfigPersonCMDInjection(db *sql.DB) domain.PersonCMD {
	personUseCase := getPersonUseCase(db)
	personCMD := personcmd.NewPersonCMD(personUseCase)

	return personCMD
}
