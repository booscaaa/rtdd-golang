package accountusecase

import "github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"

func (usecase accountUsecase) GetByID(id int32) (*domain.Account, error) {
	account, err := usecase.repository.GetByID(id)

	if err != nil {
		return nil, err
	}

	return account, nil
}
