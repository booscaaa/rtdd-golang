package domain

import (
	"fmt"
	"net/http"
)

type Person struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func (person Person) isValid() error {
	if person.Name == "" {
		return fmt.Errorf("the name is empty")
	}

	return nil
}

func NewPerson(id int32, name string, age int32) (*Person, error) {
	person := Person{
		ID:   id,
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
	Fetch(response http.ResponseWriter, request *http.Request)
	GetByID(response http.ResponseWriter, request *http.Request)
}

type PersonCMD interface {
	Fetch()
	GetByID(id int)
}

type PersonUseCase interface {
	Fetch() (*[]Person, error)
	GetByID(id int) (*Person, error)
}

type PersonRepository interface {
	Fetch() (*[]Person, error)
	GetByID(id int) (*Person, error)
}