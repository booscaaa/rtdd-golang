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
func NewPersonService(usecase domain.PersonUseCase) domain.PersonServiceServer {
	return &service{
		usecase: usecase,
	}
}

func (service service) Fetch(
	context context.Context,
	request *domain.FetchPersonRequest,
) (*domain.FetchPersonResponse, error) {
	people, err := service.usecase.Fetch()

	if err != nil {
		return nil, err
	}

	return &domain.FetchPersonResponse{People: people}, nil
}

func (service service) GetByID(
	context context.Context,
	request *domain.GetByIDRequest,
) (*domain.GetByIDResponse, error) {
	person, err := service.usecase.GetByID(request.Id)

	if err != nil {
		return nil, err
	}

	return &domain.GetByIDResponse{Person: person}, nil
}
