package accountrepository

import (
	"context"
	"fmt"
	"time"

	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
	"github.com/jackc/pgx/v4"
)

func (db accountRepository) Update(
	id int32,
	login string,
	name string,
	email string,
	accountGroupID int32,
	active bool,
) (*domain.Account, error) {
	var idA int32
	var loginA string
	var nameA string
	var passwordA string
	var emailA string
	var accountGroupIDA int32
	var createdDateA time.Time
	var activeA bool

	ctx := context.Background()

	err := db.databaseConnection.QueryRow(
		ctx,
		`UPDATE account SET 
			login = $1, 
			email = $2, 
			account_group_id = $3, 
			active = $4,
			name = $5
			WHERE id = $6
		returning *;`,
		login, email, accountGroupID, active, name, id,
	).Scan(
		&idA,
		&loginA,
		&emailA,
		&nameA,
		&passwordA,
		&createdDateA,
		&accountGroupIDA,
		&activeA,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("Account not updated")
	}

	if err != nil {
		return nil, err
	}

	account, err := domain.NewAccount(
		idA,
		loginA,
		nameA,
		passwordA,
		emailA,
		createdDateA,
		accountGroupIDA,
		activeA,
	)

	if err != nil {
		return nil, err
	}

	return account, nil
}
