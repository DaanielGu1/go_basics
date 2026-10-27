package database

import (
	"Treino-em-go/models"
	"Treino-em-go/tools"
)

func GetUserById(db *DB, id int) *models.User { // OK
	current := db.users.GetRear()
	for range db.users.GetLength() {
		if current.GetData().ID == id {
			return current.GetData()
		}

		current.GetNext()
	}
	return nil
}

func GetUserByName(db *DB, name string) *models.User { // OK
	current := db.users.GetRear()

	for range db.users.GetLength() {
		if current.GetData().Name == name {
			return current.GetData()
		}

		current = current.GetNext()
	}
	return nil
}

func CreateUserFromTerminal(db *DB, name string, password string, groupsStr ...string) bool { // OK
	if !GroupsExists(db, groupsStr) {
		return false
	}
	groups := ListGroupsByName(db, groupsStr)
	id := db.users.GetLength()
	if groups == nil {
		groups = []models.Group{}
	}

	user := models.CreateNewUser(id+1, name, groups, password)
	db.users.Add(user)

	return true
}

func (db *DB) RemoveUser(index int) *models.User { // OK
	user := db.users.Remove(index)

	return user
}

func (db *DB) FindUser(searcher tools.ISearcher[*models.User]) {
	db.users.Find(searcher)
}
