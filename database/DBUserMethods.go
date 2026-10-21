package database

import (
	"Treino-em-go/models"
	"Treino-em-go/tools"
)

func GetUserByID(db *DB, id int) *models.User {
}

func GetUserByName(db *DB, name string) *models.User {
	for _, user := range db.users {
		if user.Name == name {
			return &user
		}
	}

	return nil
}

func CreateUserFromTerminal(db *DB, name string, password string, groupsStr ...string) bool {
	if !GroupExists(db, groupsStr) {
		return false
	}
	groups := ListGroupsByName(db, groupsStr)
	id := db.users.Lenght()

	user := models.NewUser(id+1, name, groups, password)
	db.users.Add(user)

	return true
}