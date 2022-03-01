package personcmd

import (
	"fmt"
	"log"
)

func (cmd cmd) Fetch() {
	people, err := cmd.usecase.Fetch()

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, person := range *people {
		fmt.Println(
			person.Name,
		)
	}
}
