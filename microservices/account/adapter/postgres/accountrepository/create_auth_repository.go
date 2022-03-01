package accountrepository

import (
	"context"
	"fmt"
	"time"

	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
	"github.com/jackc/pgx/v4"
)

func (db accountRepository) CreateAuth(
	typeToken string,
	hash string,
	token string,
	accountID int32,
	revoked bool,
) (*domain.Auth, error) {
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
		`INSERT INTO auth (type, hash, token, account_id, revoked) VALUES ($1, $2, $3, $4, $5) returning *;`,
		typeToken,
		hash,
		token,
		accountID,
		revoked,
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
		return nil, fmt.Errorf("Auth not created")
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
