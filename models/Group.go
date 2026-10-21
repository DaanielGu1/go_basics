package models

type Group struct {
	ID   int
	Name string
}

func NewGroup(id int, group_name string) *Group {
	/*
	Cria o modelo struct da criação de um novo grupo no banco de dados, que sera utilizado
	no "CreateGroup"
	*/
	return &Group{ID: id, Name: group_name}
}
