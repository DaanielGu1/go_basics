package in

import (
	"Treino-em-go/database"
	"fmt"
	"io"
)

func NewUser(r io.Reader, w io.Writer) error{ // OK
	fmt.Println("Creating a new user...")
	name, err := IOGetString("Type user's name: ", r, w)
	
	if err != nil {
		return err
	}

	password, err := IOGetString("Type user's password: ", r, w)
	if err != nil {
		return err
	}
	for true {
		passwordCheck, err := IOGetString("Type user's password again: ", r, w)
		if err != nil {
			return err
		}
		if passwordCheck != password {
			fmt.Println("Wrong password!")
			continue
		} else {
			database.CreateUser(database.Database, name, password)
		}
		return err
	}
	return err
}

func PrintUser(r io.Reader, w io.Writer) error { // OK
	name, err := IOGetString("Type wanted user's name: ", r, w)
	if err != nil {
		return err
	}
	user := database.GetUserByName(database.Database, name)

	fmt.Printf("User's ID: %d, User's name: %s\n", user.ID, user.Name)

	return err
}
