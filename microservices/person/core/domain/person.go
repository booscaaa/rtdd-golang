package domain

import (
	"fmt"
)

func (person Person) isValid() error {
	if person.Name == "" {
		return fmt.Errorf("the name is empty")
	}

	return nil
}

func NewPerson(id int32, name string, age int32) (*Person, error) {
	person := Person{
		Id:   id,
		Age:  age,
		Name: name,
	}

	err := person.isValid()

	if err != nil {
		return nil, err
	}

	return &person, nil
}

type PersonUseCase interface {
	Fetch() ([]*Person, error)
	GetByID(id int32) (*Person, error)
}

type PersonRepository interface {
	Fetch() ([]*Person, error)
	GetByID(id int32) (*Person, error)
}
