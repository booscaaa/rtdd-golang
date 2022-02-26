package person_repository

import (
	"database/sql"

	"github.com/booscaaa/rtdd-golang/structs/core/domain"
)

type repository struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) domain.PersonRepository {
	return &repository{
		db: db,
	}
}
