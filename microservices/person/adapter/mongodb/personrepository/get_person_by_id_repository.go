package personrepository

import (
	"context"
	"fmt"

	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (repository repository) GetByID(id int32) (*domain.Person, error) {
	ctx := context.TODO()
	collection := repository.db.Collection("person")

	filter := bson.M{"id": id}

	fmt.Println(`bson.M{“_id”: docID}:`, filter)

	var person domain.Person

	if err := collection.FindOne(ctx, filter).Decode(&person); err != nil {
		return nil, err
	}

	return &person, nil
}
