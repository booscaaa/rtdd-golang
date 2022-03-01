package accountrepository

import (
	"context"
	"fmt"
	"time"

	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
	"github.com/jackc/pgx/v4"
)

func (db accountRepository) Delete(id int32) (*domain.Account, error) {
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
		`DELETE FROM account where id = $1 returning *;`,
		id,
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
		return nil, fmt.Errorf("Account not deleted")
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
