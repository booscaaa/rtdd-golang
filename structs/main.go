package main

import (
	"fmt"

	"github.com/booscaaa/structs/core/domain"
)

func main() {
	person, err := domain.NewPerson(1, "", 24)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(person)
}
