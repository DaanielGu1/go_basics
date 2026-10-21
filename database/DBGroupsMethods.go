package database

import (
	"Treino-em-go/models"
	"slices"
	"Treino-em-go/tools"
)

func GroupExists(db *DB, group models.Group) bool {
	if slices.Contains(db.groups, group) {
		return true
	}

	return false
}

func ListGroupsByName(db *DB, names []string) []models.Group {
	grouplist := make([]models.Group,len(names))
	current := db.groups.re
	for range db.groups.Lenght() {
		for _,name := range names {
			if (current.GetData().Name ==name) {
				grouplist = append(grouplist, *current.GetData())
			}
			current = current.GetNext()
		}
	}
	return grouplist
}
