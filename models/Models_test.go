package models

import (
	"testing"
)

func TestUsers(t *testing.T)  {
	group:= CreateNewGroup(1, "group1")
	otherGroup:= CreateNewGroup(2, "group2")
	groupList := []Group{*group}
	
	user1 := CreateNewUser(3,"user1",groupList,"password1")
	if user1 == nil {
		t.Errorf("CreateNewUser() == nil; function can't create a user in memory")
	}
	
	returnedGroup := user1.GetGroup("group1")
	if returnedGroup != group {
		t.Errorf("GetGroup() = nil; want %p; funcion isn't able to return correctly", group)
	}
	returnedGroup = user1.GetGroup("noGroup")
	if returnedGroup != nil {
		t.Errorf("GetGroup() = %p; want nil; function is returning a group that wasn't required", returnedGroup)
	}

	if !user1.InGroup("group1"){
		t.Errorf("User.InGroup() = false; want true; function can't recognize the group in the user")
	} 

	if user1.InGroup("noGroup"){
		t.Errorf("User.InGroup()= true; want false; function can't recognize that the user isn't in the group")
	}

	if user1.AddGroup(group){
		t.Errorf("User.AddGroup() = true; want false; function is adding a group that is already in the user")
	}

	if user1.AddGroup(nil) {
		t.Errorf("User.AddGroup() = true; want false; function is adding nil to user's groups")
	}

	if !user1.AddGroup(otherGroup) {
		t.Errorf("User.AddGroup() = false; want true; function isn't adding a group that isn't in the user list of groups")
	}

	/* returnedGroup = user1.RemoveGroupFromUser(group) 
	if returnedGroup == nil {
		t.Errorf("User.RemoveGroupFromUser() = nil; want %p; function should return the removed group", group)
	}

	returnedGroup = user1.RemoveGroupFromUser(CreateNewGroup(3,"groupName"))
	if returnedGroup != nil {
		t.Errorf("User.RemoveGroupFromUser() = %p; want nil; function shouldn't return a ptr when the group isn't in the user's groups", returnedGroup)
	} */
}