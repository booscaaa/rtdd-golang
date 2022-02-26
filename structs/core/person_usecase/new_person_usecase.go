package person_usecase

import "github.com/booscaaa/rtdd-golang/structs/core/domain"

type usecase struct {
	repository domain.PersonRepository
}

func NewPersonUseCase(repository domain.PersonRepository) domain.PersonUseCase {
	return &usecase{
		repository: repository,
	}
}
