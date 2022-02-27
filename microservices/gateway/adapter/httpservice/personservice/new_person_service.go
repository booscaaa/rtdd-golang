package personservice

import "github.com/booscaaa/rtdd-golang/microservices/gateway/core/domain"

type service struct {
	usecase domain.PersonUseCase
}

// NewPersonService .
func NewPersonService(usecase domain.PersonUseCase) domain.PersonService {
	return &service{
		usecase: usecase,
	}
}