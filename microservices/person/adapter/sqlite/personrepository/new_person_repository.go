package personrepository

import (
	"database/sql"

	"github.com/booscaaa/rtdd-golang/microservices/person/core/domain"
)

type repository struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) domain.PersonRepository {
	return &repository{
		db: db,
	}
}
