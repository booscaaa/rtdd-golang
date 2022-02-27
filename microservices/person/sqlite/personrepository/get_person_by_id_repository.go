package personrepository

import "github.com/booscaaa/rtdd-golang/microservices/gateway/core/domain"

func (repository repository) GetByID(id int) (*domain.Person, error) {
	var idA int32
	var nameA string
	var ageA int32

	err := repository.db.
		QueryRow("SELECT id, name, age FROM person WHERE id = ?", id).
		Scan(&idA, &nameA, &ageA)

	if err != nil {
		return nil, err
	}

	person, _ := domain.NewPerson(idA, nameA, ageA)

	return person, nil
}
