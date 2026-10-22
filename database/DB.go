// Ok

package database

import (
	"Treino-em-go/tools"
	"Treino-em-go/models"
)

type DB struct {
	users  *tools.SLL[*models.User]
	groups *tools.SLL[*models.Group]
}

var Database = &DB{
	users:  &tools.SLL[*models.User]{},
	groups: &tools.SLL[*models.Group]{},
}
