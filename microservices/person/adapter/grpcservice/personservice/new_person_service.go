package personservice

import (
	"context"

	"github.com/booscaaa/rtdd-golang/microservices/person/core/domain"
)

type service struct {
	usecase domain.PersonUseCase
	domain.UnimplementedPersonServiceServer
}

// NewPersonService .
func NewPersonService(usecase domain.PersonUseCase) domain.PersonService {
	return &service{
		usecase: usecase,
	}
}

func (service service) Fetch(
	context context.Context,
	request *domain.FetchRequest,
) (*domain.FetchResponse, error) {
	people, err := service.usecase.Fetch()

	if err != nil {
		return nil, err
	}

	return &domain.FetchResponse{People: people}, nil
}
