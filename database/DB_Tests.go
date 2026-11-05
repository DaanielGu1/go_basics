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

// ---- User's Tests ---- \\

func TestCreateUser(t *testing.T) {
	CreateGroupFromTerminal(Database, "Group1")

	CreateUserFromTerminal(Database, "JhonDoe", "123", "Group1")

	if !CreateUserFromTerminal(Database, "JhonDoe", "123", "Group1") {
		t.Errorf("CreateUserFromTerminal(); There's a problem creating a user with that group(s)")
	}

	if Database.users.Get(0) == nil {
		t.Errorf("CreateUserFromTerminal(); The user wasn't created.")
	}
}

func TestGetUser(t *testing.T) {
	if GetUserById(Database, 1) != nil {
		t.Errorf("GetUserbyID(); The function found a non-existing user.")
	}

	if GetUserByName(Database, "JhonDoe") != nil {
		t.Errorf("GetUserbyName(); The function a non-existing user.")
	}

	CreateUserFromTerminal(Database, "JhonDoe", "123")
	CreateUserFromTerminal(Database, "JaneDoe", "123")
	CreateUserFromTerminal(Database, "Heisemberg", "123")

	if GetUserById(Database, 2) != Database.users.Get(1) {
		t.Errorf("GetUserByID(); The function found a wrong user.")
	}

	if GetUserByName(Database, "JhonDoe") != Database.users.Get(0) {
		t.Errorf("GetUserByName(); The function found a wrong user.")
	}
}

func TestRemoveUser(t *testing.T) {
	CreateUserFromTerminal(Database, "JhonDoe", "123")
	CreateUserFromTerminal(Database, "JaneDoe", "123")
	CreateUserFromTerminal(Database, "Heisemberg", "123")

	user := Database.users.Get(1)

	removed := Database.RemoveUser(1)

	if Database.users.GetLength() != 2 {
		t.Errorf("RemoveUser(); The user wasn't removed from database.")
	}

	if user != removed {
		t.Errorf("RemoveUser(); The function removed the wrong user.")
	}
}
