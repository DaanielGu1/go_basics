package in

import (
	"Treino-em-go/database"
	"fmt"
	"io"
)

func NewUser(r io.Reader, w io.Writer) error{ // OK
	fmt.Fprint(w, "Creating a new user...")
	name, err := IOGetString("Type user's name: ", r, w)
	
	if err != nil {
		return err
	}

	password, err := IOGetString("Type user's password: ", r, w)
	
	if err != nil {
		return err
	}
	
	for {
		passwordCheck, err := IOGetString("Type user's password again: ", r, w)
		if err != nil {
			return err
		}

		if passwordCheck != password {
			fmt.Fprint(w, "Wrong password!")
			continue
		} else {
			database.CreateUser(database.Database, name, password)
		}
		
		return err
	}
}

func PrintUser(r io.Reader, w io.Writer) error { // OK
	name, err := IOGetString("Type wanted user's name: ", r, w)
	if err != nil {
		return err
	}

	user := database.GetUserByName(database.Database, name)
	
	if user != nil {
		fmt.Fprintf(w, "User's ID: %d, User's name: %s\n", user.ID, user.Name)
	} else {
		fmt.Fprint(w, "Couldn't find this user.")
	}

	return err
}
