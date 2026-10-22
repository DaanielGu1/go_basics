package main

import (
	"Treino-em-go/database"
	"Treino-em-go/in"
)

func main() {
	
		in.NewUser() // Creating a new user correctly
		in.NewUser()
		in.NewUser()
	
		/* in.PrintUser() // Printing a user from database correctly

		in.NewGroup() // Creating a new group correctly
		in.NewGroup()

		in.PrintGroup() // Printing a group from database correctly */

		database.Database.RemoveUser(2) // Deleting a user from database correctly

		in.PrintUser()
		in.PrintUser()
		
}
