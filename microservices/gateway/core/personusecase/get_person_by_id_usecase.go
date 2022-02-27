package personusecase

import "github.com/booscaaa/rtdd-golang/microservices/gateway/core/domain"

func (usecase usecase) GetByID(id int) (*domain.Person, error) {
	person, err := usecase.repository.GetByID(id)

	if err != nil {
		return nil, err
	}

	return person, nil
}
