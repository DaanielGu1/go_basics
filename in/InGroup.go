package in

import (
	"Treino-em-go/database"
	"fmt"
	"io"
)

func NewGroup(r io.Reader, w io.Writer) error { // OK
	fmt.Fprint(w, "Creating a new group...")
	name, err := IOGetString("Type group's name: ", r, w)

	if err != nil {
		return err
	}

	ok := database.CreateGroup(database.Database, name)

	if !ok {
		fmt.Fprint(w, "This group already exists.")
	}
	
	return err
}