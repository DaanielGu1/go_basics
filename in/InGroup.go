package in

import (
	"Treino-em-go/database"
	"fmt"
)

func NewGroup() error {
	fmt.Println("Creating a new group...")
	name, err := IOGetString("Type group's name: ")

	if err != nil {
		return err
	}

	ok := database.CreateGroupFromTerminal(database.Database, name)

	if !ok {
		fmt.Println("This group already exists.")
	}
	
	return err
}

func PrintGroup() error {
	name, err := IOGetString("Type wanted group's name: ")
	
	if err != nil {
		return err
	}

	group := database.GetGroupByName(database.Database, name)

	fmt.Printf("Group's ID: %d, Group's Name: %s\n", group.ID, group.Name)

	return  err
}