package personusecase

import "github.com/booscaaa/rtdd-golang/microservices/person/core/domain"

type usecase struct {
	repository domain.PersonRepository
}

// NewPersonUseCase .
func NewPersonUseCase(repository domain.PersonRepository) domain.PersonUseCase {
	return &usecase{
		repository: repository,
	}
}
