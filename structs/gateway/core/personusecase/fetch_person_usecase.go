package personusecase

import "github.com/booscaaa/rtdd-golang/gateway/core/domain"

func (usecase usecase) Fetch() (*[]domain.Person, error) {
	people, err := usecase.repository.Fetch()

	if err != nil {
		return nil, err
	}

	return people, nil
}
