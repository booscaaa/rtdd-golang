package accountrepository

import (
	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type accountRepository struct {
	databaseConnection *pgxpool.Pool
}

func NewAccountRepository(connection *pgxpool.Pool) domain.AccountRepository {
	return &accountRepository{
		databaseConnection: connection,
	}
}
