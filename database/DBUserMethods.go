package database

import (
	"Treino-em-go/models"
	"Treino-em-go/tools"
)

func GetUserById(db *DB, id int) *models.User { // OK
	for i := range db.users.GetLength() {
		current := db.users.Get(i)
	
		if current.ID == id {
			return current
		}
	}
	return nil
}

func GetUserByName(db *DB, name string) *models.User { // OK
	for i := range db.users.GetLength() {
		current := db.users.Get(i)
	
		if current.Name == name {
			return current
		}
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
