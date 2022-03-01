package accountusecase

import "github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"

func (usecase accountUsecase) GetByName(name string) (*domain.Account, error) {
	account, err := usecase.repository.GetByName(name)

	if err != nil {
		return nil, err
	}

	return account, nil
}
