package accountrepository

import (
	"context"
	"fmt"
	"time"

	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
	"github.com/jackc/pgx/v4"
)

func (db accountRepository) DeleteAuth(id int32) (*domain.Auth, error) {
	var idA int32
	var typeTokenA string
	var hashA string
	var tokenA string
	var accountIDA int32
	var revokedA bool
	var createdDateA time.Time

	ctx := context.Background()

	err := db.databaseConnection.QueryRow(
		ctx,
		`DELETE FROM auth where account_id = $1 returning *;`,
		id,
	).Scan(
		&idA,
		&typeTokenA,
		&hashA,
		&tokenA,
		&accountIDA,
		&revokedA,
		&createdDateA,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("Account auth not deleted")
	}

	if err != nil {
		return nil, err
	}

	auth, err := domain.NewAuth(idA, typeTokenA, hashA, tokenA, accountIDA, revokedA, createdDateA)

	if err != nil {
		return nil, err
	}

	return auth, nil
}
