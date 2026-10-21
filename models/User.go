package models

type User struct {
	ID     int
	Name   string
	psw    string
	Groups []Group
}

func NewUser(id int, name string, password string, groups []Group) *User {
	/*
		Essa função vai definir a estrutura que será usada na criação de usuário
		na função "CreateUser" do banco de dados.
	*/

	return &User{ID: id, Name: name, psw: password, Groups: groups}
}

func (user *User) GetGroup(groupName string) *Group {
	/*
		Esse método vai expor se um grupo específico está dentro da lista de grupos
		do usuário em questão.
	*/

	for _, group := range user.Groups {
		if group.Name == groupName {
			return &group
		}
	}

	return nil
}

func (user *User) InGroup(groupName string) bool {
	/*
		Esse método vai expor se um usuário específico está dentro de uma lista de usuários
		de um grupo.
	*/

	for _, group := range user.Groups {
		if group.Name == groupName {
			return true
		}
	}

	return false
}

func (user *User) AddGroup(group Group) bool {
	// Esse método adiciona um novo grupo à lista de de grupos de um determinado usuário.

	if user.InGroup(group.Name) {
		return false
	}

	user.Groups = append(user.Groups, group)

	return true
}

func (user *User) RemoveGroup(groupName string) *Group {
	/*
		Esse método serve para remover um determinado grupo da lista de grupos de um único
		usuário.
	*/

	for i, group := range user.Groups { // Usando um for para percorrer a lista de grupos de um usuário.
		if group.Name == groupName { // Verifica se o nome do grupo atual é igual ao do parâmetro.
			user.Groups = append(user.Groups[:i], user.Groups[i+1:]...) // Copiar a lista até antes do grupo atual e depois do grupo atual.
			return &group
		}
	}
	return nil // Caso ele não encontre o grupo na lista, retorna vazio.
}
