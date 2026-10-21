// Ok

package database

import (
	"Treino-em-go/models"
	"Treino-em-go/tools"
)

type DB struct {
	users  *tools.SLL[*models.User]
	groups *tools.SLL[*models.Group]
}

var Database = &DB{
	users: &tools.SLL[*models.User]{},
	groups: &tools.SLL[*models.Group]{},
}

