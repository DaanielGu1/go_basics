package database

import (
	"Treino-em-go/models"
	"Treino-em-go/tools"
)

func GroupsExists(db *DB, names []string) bool {
	var cont int
	
	for i := range db.groups.GetLength() {
		current := db.groups.Get(i)
	
		for index := range names {
			if current.Name == names[index] {
				cont += 1
			}
		}
	}
	if cont == len(names) {
		return true
	}
	return false
}

func ListGroupsByName(db *DB, names []string) []models.Group {
	grouplist := make([]models.Group, len(names))
	
	for i := range db.users.GetLength() {
		current := db.groups.Get(i)

		for _, name := range names {
			if current.Name == name {
				grouplist = append(grouplist, *current)
			}
		}
	}
	return grouplist
}

func GetGroupByName(db *DB, name string) *models.Group {
	for i := range db.groups.GetLength() {
		current := db.groups.Get(i)
	
		if current.Name == name {
			return current
		}
	}
	return nil
}

func CreateGroupFromTerminal(db *DB, name string) bool {
	groupL := []string{name}

	if GroupsExists(db, groupL) {
		return false
	}

	id := db.groups.GetLength() + 1
	group := models.CreateNewGroup(id, name)
	db.groups.Add(group)

	return true
}

func (db *DB) RemoveGroup(id int) *models.Group {
	group := db.groups.Remove(id)

	return group
}

func (db *DB) Find(searcher tools.ISearcher[*models.Group]) {
	db.groups.Find(searcher)
}
