package personrepository

import "github.com/booscaaa/rtdd-golang/microservices/person/core/domain"

func (repository repository) Fetch() ([]*domain.Person, error) {
	people := []*domain.Person{}

	rows, err := repository.db.Query("SELECT id, name, age FROM person")
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var idA int32
		var nameA string
		var ageA int32

		err := rows.Scan(&idA, &nameA, &ageA)

		if err != nil {
			return nil, err
		}

		person, _ := domain.NewPerson(idA, nameA, ageA)

		people = append(people, person)
	}

	return people, nil
}
