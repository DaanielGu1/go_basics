package in

import (
	"fmt"
	"Treino-em-go/database"
)

func NewUser() error {
	fmt.Println("Creating a new user...")
	name, err := IOGetString("Type user's name")

	if err != nil {
		return err
	}

	password, err := IOGetString("Type user's password")
	if err != nil {
		return err
	}
	for true {
		passwordCheck, err := IOGetString("Type user's password again")
		if err != nil {
			return err
		}
		if passwordCheck != password {
			fmt.Println("Wrong password!")
			continue
		} else {
			database.CreateUserFromTerminal(database.Database, name, password)
		}
		return err
	}
	return err
}

func PrintUser() error {
	name, err := IOGetString("Type wanted user's name")
	if err != nil {
		return err
	}
	user := database.GetUserByName(database.Database, name)

	fmt.Printf("User ID: %d, User name: %s", user.ID, user.Name)

	return err
}
