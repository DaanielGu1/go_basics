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

func (user *User) RemoveGroupFromUser(group *Group) *Group {
	for index, userGroup := range user.Groups {
		if userGroup == *group {
			user.Groups = append(user.Groups[:index], user.Groups[index+1:]...)
			return group
		}
	}
	return nil
}
