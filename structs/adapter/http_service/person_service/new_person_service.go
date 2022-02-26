package person_service

import "github.com/booscaaa/rtdd-golang/structs/core/domain"

type service struct {
	usecase domain.PersonUseCase
}

func NewPersonService(usecase domain.PersonUseCase) domain.PersonService {
	return &service{
		usecase: usecase,
	}
}
