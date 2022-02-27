package personcmd

import (
	"fmt"
	"log"
)

func (cmd cmd) GetByID(
	id int,
) {
	person, err := cmd.usecase.GetByID(
		id,
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(
		person.Name,
	)
}
