package accountusecase

import "github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"

func (usecase accountUsecase) Create(c *domain.Account) (*domain.Account, error) {
	account, err := usecase.repository.Create(
		c.Login,
		c.Name,
		c.Password,
		c.Email,
		c.AccountGroupID,
		c.Active,
	)

	if err != nil {
		return nil, err
	}

	return account, nil
}
