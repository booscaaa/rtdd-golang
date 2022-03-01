package accountservice

import (
	"context"
	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
)

type service struct {
	usecase domain.AccountUsecase
	domain.UnimplementedAccountServiceServer
}

// NewAccountService .
func NewAccountService(usecase domain.AccountUsecase) domain.AccountServiceServer {
	return &service{
		usecase: usecase,
	}
}

func (service service) Login(
	context context.Context,
	request *domain.LoginAccountRequest,
) (*domain.LoginAccountResponse, error) {
	jwt, err := service.usecase.VerifyAuthentication(request.User, request.Password)

	if err != nil {
		return nil, err
	}

	return jwt, nil
}

// func (service service) GetByID(
// 	context context.Context,
// 	request *domain.GetByIDRequest,
// ) (*domain.GetByIDResponse, error) {
// 	person, err := service.usecase.GetByID(request.Id)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &domain.GetByIDResponse{Person: person}, nil
// }
