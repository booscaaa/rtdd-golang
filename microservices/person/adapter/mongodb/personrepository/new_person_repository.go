package personrepository

import (
	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Database
}

func NewPersonRepository(db *mongo.Database) domain.PersonRepository {
	return &repository{
		db: db,
	}
}
