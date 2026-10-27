package database

import (
	"Treino-em-go/models"
	"Treino-em-go/tools"
)

func GroupsExists(db *DB, names []string) bool {
	var cont int
	current := db.groups.GetRear()
	for range db.groups.GetLength() {
		for index := range names {
			if current.GetData().Name == names[index] {
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
	current := db.groups.GetRear()
	for range db.groups.GetLength() {
		for _, name := range names {
			if current.GetData().Name == name {
				grouplist = append(grouplist, *current.GetData())
			}
			current = current.GetNext()
		}
	}
	return grouplist
}

func GetGroupByName(db *DB, name string) *models.Group {
	current := db.groups.GetRear()

	for range db.groups.GetLength() {
		if current.GetData().Name == name {
			return current.GetData()
		}

		current = current.GetNext()
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
