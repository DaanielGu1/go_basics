package models

type User struct {
	ID     int
	Name   string
	psw    string
	Groups []Group
}

func CreateNewUser(id int, name string, groups []Group, password string) *User {
	return &User{ID: id, Name: name, Groups: groups, psw: password}
}

func (user *User) GetGroup(groupName string) *Group {
	for _, group := range user.Groups {
		if group.Name == groupName {
			return &group
		}
	}
	return nil
}

func (user *User) InGroup(groupName string) bool {
	if user.Groups != nil {
		for _, group := range user.Groups {
			if group.Name == groupName {
				return true
			}
		}
	}
	return false
}

func (user *User) AddGroup(group *Group) bool {
	// Esse método adiciona um novo grupo à lista de de grupos de um determinado usuário.

	if group == nil {
		return false
	}

	if user.InGroup(group.Name) {
		return false
	}

	user.Groups = append(user.Groups, *group)
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
