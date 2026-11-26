package database

import (
	"Treino-em-go/tools"
	"testing"
)

func TestDatabaseAttributes(t *testing.T) {
	if Database.users == nil {
		t.Errorf("Database.users == nil; want ptr -> &tools.SLL[*models.User]")
	}
	if Database.groups == nil {
		t.Errorf("Database.groups == nil; want ptr -> &tools.SLL[*models.Group]")
	}
}

func TestGroupMethods(t *testing.T) {
	nameList := []string{"group1", "group2", "group3"}

	CreateGroup(Database, "group1")
	if CreateGroup(Database, "group1") {
		t.Errorf("CreateGroup() = true; want false; Function shouldn't create a group if a group with it's name already exists")
	}

	if Database.groups.GetLength() != 1 {
		t.Errorf("CreateGroupFromTerminal(); SLL Length = %d, want %d; function has issues to create the group", Database.groups.GetLength(), 1)
	}

	if GroupsExists(Database, nameList) {
		t.Errorf("GroupsExists() = %t; want %t; function should return false due to lacking names on name list", true, false)
	}

	group := GetGroupByName(Database, "group2")
	if group != nil {
		t.Errorf("GetGroupByName = %p; want nil; function is returning a group that is not the required", group)
	}

	CreateGroup(Database, "group2")
	CreateGroup(Database, "group3")

	group = GetGroupByName(Database, "group2")
	if group == nil {
		t.Errorf("GetGroupByName = nil; want ptr to group; function isn't returning a group even if the required group exists in the list")
	}

	if !GroupsExists(Database, nameList) {
		t.Errorf("GroupExists() = %t; want %t; function has issues to identify the wanted groups in the list", false, true)
	}

	removedGroup := Database.RemoveGroup(2)
	if removedGroup == nil {
		t.Errorf("RemoveGroup() = nil; want %p; function should return a pointer to a group when removing it", group)
	}

	if Database.groups.GetLength() != 2 {
		t.Errorf(`RemoveGroup(); SLL Length = %d ; want %d Length ain't getting 
		reduced when removing a group`, Database.groups.GetLength(), 2)
	}

	group = Database.groups.Get(0)

	searcher := tools.NewFindFirstSearcher(tools.NewGroupComparable(group))

	Database.FindGroup(searcher)
}

func TestUserMethods(t *testing.T) {
	CreateGroup(Database, "firstUserGroup")
	CreateGroup(Database, "secondUserGroup")

	if CreateUser(Database, "user1", "psw1", "userGroup") {
		t.Errorf("CreateUserFromTerminal() = %t; want %t; user is being created with a non existent group", true, false)
	}

	if !CreateUser(Database, "user1", "psw1", "firstUserGroup", "secondUserGroup") {
		t.Errorf("CreateUserFromTerminal() = %t; want %t; function has issues when trying to create a user with groups", false, true)
	}

	user := GetUserByName(Database, "user1")

	if user == nil {
		t.Errorf("GetUserByName() = nil; want ptr to user; function has issues to get the user by it's name;")
	}

	user = GetUserByName(Database, "user100")

	if user != nil {
		t.Errorf("GetUserByName() = nil; want ptr to user; function has issues to get the user by it's name;")
	}

	user = GetUserById(Database, 1)

	if user == nil {
		t.Errorf("GetUserById() = nil; want ptr to user; function has issues to get the user by it's id")
	}

	user = GetUserById(Database, 1920)

	if user != nil {
		t.Errorf("GetUserById() = %p; want nil; function shouldn't return a pointer if there no user with required id", user)
	}

	CreateUser(Database, "user2", "psw2")

	user = GetUserById(Database, 2)

	removedUser := Database.RemoveUser(1)

	if removedUser != user {
		t.Errorf("RemoveUser() = %p; want %p; function has issues to get the right user's removed index", removedUser, user)
	}

	if Database.users.GetLength() != 1 {
		t.Errorf("RemoveUser() = %d; want %d; function is not reducing list's length", Database.groups.GetLength(), 1)
	}

	Database.RemoveUser(0)

	CreateUser(Database, "userToFind", "password")
	user = GetUserById(Database, 1)
	searcher := tools.NewFindFirstSearcher(tools.NewUserComparable(user))
	Database.FindUser(searcher)
	if searcher.GetIndex() != 0 {
		t.Errorf("searcher.FindUser(); searcher.GetIndex() = %d; want 0; method is not being able to execute Find() correctly", searcher.GetIndex())
	}
}