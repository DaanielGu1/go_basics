package database

import (
	"testing"
)

// ---- Var DB test ---- \\

func TestDBLists(t *testing.T) {
	if Database.users != nil {
		t.Errorf("Database.users; The list is not initialized empty")
	}

	if Database.groups != nil {
		t.Errorf("Database.groups; The list is not initialized empty")
	}
}

// ---- Group's Tests ---- \\

func TestGroupsMethods(t *testing.T) {
	nameList:= []string{"group1", "group2", "group3"}
	
	shouldBeEmpty := ListGroupsByName(Database, nameList)
	
	if len(shouldBeEmpty) != 0 {
		t.Errorf("len(ListGroupsByName()) = %d; want %d; When used on an empty SLL, len(ListGroupsByName()) should be 0", len(shouldBeEmpty), 0)
	}

	CreateGroupFromTerminal(Database, "group1")

	if Database.groups.GetLength() != 1 {
		t.Errorf("CreateGroupFromTerminal(); SLL Length = %d, want %d; Function has issues to create the user",Database.groups.GetLength(), 1)
	}

	if GroupsExists(Database, nameList) {
		t.Errorf("GroupsExists() = %t; want %t; Function should return false due to lacking names on name list", true, false)
	}

	group:= GetGroupByName(Database, "group2")
	if group != nil {
		t.Errorf("GetGroupByName = %p; want nil; Function is returning a group that is not the required", group)
	}

	CreateGroupFromTerminal(Database, "group2")
	CreateGroupFromTerminal(Database, "group3")

	group = GetGroupByName(Database, "group2")
	if group == nil {
		t.Errorf("GetGroupByName = nil; want %p; Function isn't returning a group even if the required group exists in the list", group)
	}

	if !GroupsExists(Database, nameList) {
		t.Errorf("GroupExists() = %t; want %t; Function has issues to identify the wanted groups in the list", false, true)
	}

	removedGroup := Database.RemoveGroup(2)
	if removedGroup == nil {
		t.Errorf("RemoveGroup() = nil; want %p; Function should return a pointer to a group when removing it", group)
	}

	if Database.groups.GetLength() != 2 {
		t.Errorf("RemoveGroup(); SLL Length = %d ; want %d Length ain't getting reduced when removing a group", Database.groups.GetLength(), 2)
	}
}

// ---- User's Tests ---- \\

func TestUserMethods(t *testing.T) {
	CreateGroupFromTerminal(Database, "Group1")
	
	CreateUserFromTerminal(Database, "JhonDoe", "123", "Group1")

	if Database.users.GetLength() != 1 {
		t.Errorf("CreateUserFromTerminal(); SLL Length = %d, want %d; Function has issues to create the user",Database.users.GetLength(), 1)
	}

	user := GetUserByName(Database, "JhonDoe")

	if user != nil {
		t.Errorf("GetUserByName = %p; want nil; Function is returning a user that is not the required", user)
	}

	CreateUserFromTerminal(Database, "JaneDoe", "123")
	CreateUserFromTerminal(Database, "Heisemberg", "123")

	user = GetUserByName(Database, "JaneDoe")
	
	if user == nil {
		t.Errorf("GetGroupByName = nil; want %p; Function isn't returning a group even if the required group exists in the list", user)
	}

	removedUser := Database.RemoveUser(2)
	
	if removedUser == nil {
		t.Errorf("RemoveUser() = nil; want %p; Function should return a pointer to a user when removing it", user)
	}

	if Database.users.GetLength() != 2 {
		t.Errorf("RemoveUser(); SLL Length = %d ; want %d Length ain't getting reduced when removing a user", Database.users.GetLength(), 2)
	}
}

func TestGetUser(t *testing.T) {
	if GetUserById(Database, 1) != nil {
		t.Errorf("GetUserbyID() = %p; want nil; The function found a non-existing user.", GetUserById(Database, 1))
	}

	if GetUserByName(Database, "JhonDoe") != nil {
		t.Errorf("GetUserbyName() = %p; want nil; The function a non-existing user.", GetUserByName(Database, "JhonDoe"))
	}

	CreateUserFromTerminal(Database, "JhonDoe", "123")
	CreateUserFromTerminal(Database, "JaneDoe", "123")
	CreateUserFromTerminal(Database, "Heisemberg", "123")

	if GetUserById(Database, 2) != Database.users.Get(1) {
		t.Errorf("GetUserByID() = %p; want %p; The function found a wrong user.", GetUserById(Database, 2), Database.users.Get(1))
	}

	if GetUserByName(Database, "JhonDoe") != Database.users.Get(0) {
		t.Errorf("GetUserByName() = %p; want %p; The function found a wrong user.", GetUserByName(Database, "JhonDoe"), Database.users.Get(0))
	}
}
