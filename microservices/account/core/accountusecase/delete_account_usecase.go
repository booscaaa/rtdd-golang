package accountusecase

import "github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"

func (usecase accountUsecase) Delete(id int32) (*domain.Account, error) {
	_, err := usecase.repository.DeleteAuth(id)

	if err != nil {
		return nil, err
	}

	account, err := usecase.repository.Delete(id)

	if err != nil {
		return nil, err
	}

	return account, nil
}
