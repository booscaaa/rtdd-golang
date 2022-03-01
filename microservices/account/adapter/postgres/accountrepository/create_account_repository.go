package accountrepository

import (
	"context"
	"fmt"
	"time"

	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
	"github.com/jackc/pgx/v4"
)

func (db accountRepository) Create(
	login string,
	name string,
	password string,
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
		`INSERT INTO account (login, name, password, email, account_group_id, active) VALUES ($1, $2, crypt($3, gen_salt('bf', 8)), $4, $5, $6, $7, $8) returning *;`,
		login,
		name,
		password,
		email,
		accountGroupID,
		active,
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
		return nil, fmt.Errorf("Account not created")
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
