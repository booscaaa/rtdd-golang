package personrepository

import (
	"context"

	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (repository repository) Fetch() ([]*domain.Person, error) {
	ctx := context.TODO()
	people := []*domain.Person{}

	collection := repository.db.Collection("person")

	rows, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for rows.Next(ctx) {
		var person domain.Person

		err := rows.Decode(&person)

		if err != nil {
			return nil, err
		}

		people = append(people, &person)
	}

	return people, nil
}
