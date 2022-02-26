package person_repository

import "github.com/booscaaa/rtdd-golang/structs/core/domain"

func (repository repository) GetByID(id int) (*domain.Person, error) {
	var idA int
	var nameA string
	var ageA int

	err := repository.db.
		QueryRow("SELECT id, name, age FROM person WHERE id = ?", id).
		Scan(&idA, &nameA, &ageA)

	if err != nil {
		return nil, err
	}

	person, _ := domain.NewPerson(idA, nameA, ageA)

	return person, nil
}
