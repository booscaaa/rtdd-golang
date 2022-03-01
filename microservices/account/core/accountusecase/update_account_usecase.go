package accountusecase

import (
	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
)

func (usecase accountUsecase) Update(c *domain.Account) (*domain.Account, error) {
	accountOld, err := usecase.repository.GetByID(c.Id)

	if err != nil {
		return nil, err
	}

	if accountOld.Password != c.Password {
		account, err := usecase.repository.UpdateWithPassword(
			c.Id,
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

	account, err := usecase.repository.Update(
		c.Id,
		c.Login,
		c.Name,
		c.Email,
		c.AccountGroupID,
		c.Active,
	)

	if err != nil {
		return nil, err
	}

	return account, nil

}
