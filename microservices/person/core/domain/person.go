package domain

import (
	context "context"
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

type PersonService interface {
	Fetch(context context.Context, request *FetchRequest) (*FetchResponse, error)
}

type PersonUseCase interface {
	Fetch() ([]*Person, error)
}

type PersonRepository interface {
	Fetch() ([]*Person, error)
}
