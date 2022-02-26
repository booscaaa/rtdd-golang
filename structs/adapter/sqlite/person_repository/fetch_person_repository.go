package person_repository

import "github.com/booscaaa/rtdd-golang/structs/core/domain"

func (repository repository) Fetch() (*[]domain.Person, error) {
	people := []domain.Person{}

	rows, err := repository.db.Query("SELECT id, name, age FROM person")
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var idA int
		var nameA string
		var ageA int

		err := rows.Scan(&idA, &nameA, &ageA)

		if err != nil {
			return nil, err
		}

		person, _ := domain.NewPerson(idA, nameA, ageA)

		people = append(people, *person)
	}

	return &people, nil
}
