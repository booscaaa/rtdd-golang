package accountusecase

import "github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"

type accountUsecase struct {
	repository domain.AccountRepository
}

func NewAccountUseCase(
	repository domain.AccountRepository,
) domain.AccountUsecase {
	return &accountUsecase{
		repository: repository,
	}
}
