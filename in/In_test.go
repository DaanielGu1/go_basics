package in

import (
	"testing" 
	"strings"
)

func TestGetString(t *testing.T) {
	var r = strings.NewReader("Daniel")
	var w = &strings.Builder{}

	name, err := IOGetString("message", r, w)
	expected := "message"

	if w.String() != expected {
		t.Errorf("IOGetString(): The function should print a differnt message.")
	}

	if err != nil {
		t.Error(err)
	}

	if name != "Daniel" {
		t.Errorf("IOGetString() = %s; want Daniel: The function got the a wrong string.", name)
	}

	if name == "" {
		t.Errorf("IoGetString(): The function got an empty string.")
	}

	_, err = IOGetString("", r, w)

	if err == nil {
		t.Errorf("IOGetString(): The function shouldn't catch an empty message.")
	}

	r = strings.NewReader("")
	_, err = IOGetString("message", r, w)

	if err == nil {
		t.Errorf("IOGetString(): The functio shouldn't catch an empty string.")
	}
}

func TestGetInt(t *testing.T) {
	var r = strings.NewReader("3")
	var w = &strings.Builder{}

	num, err := IOGetInt("message", r, w)

	if err != nil {
		t.Error(err)
	}

	if num != 3 {
		t.Errorf("IOGetInt() = %d; want Daniel: The function got the a wrong string.", num)
	}

	r = strings.NewReader("!")
	num, err = IOGetInt("message", r, w)

	if err == nil {
		t.Error(err)
	}

	r = strings.NewReader("999999999999999999999999999999999999999999999999999999999999999999999999")
	num, err = IOGetInt("message", r, w)

	if err == nil {
		t.Errorf("IOGetInt() = %d; want a number out of range error.", num)
	}

	r = strings.NewReader("")
	num, err = IOGetInt("message", r, w)

	if err == nil {
		t.Errorf("IOGetInt(): The function shouldn't get an empty string.")
	}

}

func TestInGroup(t *testing.T) {
	var r = strings.NewReader("Ribamar")
	var w = &strings.Builder{}

	err := NewGroup(r, w)

	if err != nil {
		t.Error(err)
	}

	r2 := strings.NewReader("Ribamar")

	err = NewGroup(r2, w)

	if err == nil {
		t.Errorf("NewGroup(): The function is creating a group that already exists")
	}

	r3 := strings.NewReader("")

	err = NewGroup(r3, w)

	if err == nil {
		t.Error(err)
	}
}

func TestInUser(t *testing.T) {
	var r = strings.NewReader("Ribamar\nAmoFlamengo\nAmoVasco\nAmoFlamengo")
	var w = &strings.Builder{}

	err := NewUser(r, w)

	if err != nil {
		t.Error(err)
	}

	r2 := strings.NewReader("\nAmoFlamengo\nAmoVasco\nAmoFlamengo")

	err = NewUser(r2, w)

	if err == nil {
		t.Errorf("NewUser(): The function is creating a nameless user")
	}

	r3 := strings.NewReader("\n\n\nAmoFlamengo")

	err = NewUser(r3, w)

	if err == nil {
		t.Errorf("NewUser(): The function is creating a incomplete user")
	}

	r4 := strings.NewReader("")

	err = NewUser(r4, w)

	if err == nil {
		t.Errorf("NewUser(): The function is creating a incomplete user")
	}

	r5 := strings.NewReader("Ribamar")

	PrintUser(r5, w)
	expected := "User's ID: 1, User's name: Ribamar"

	if w.String() != expected {
		t.Errorf("PrintUser: The function can't print properly.")
	}

	err = PrintUser(r4, w)

	if err == nil {
		t.Errorf("PrintUser(): The function need a name.")
	}

	r6 := strings.NewReader("Cleitin")

	PrintUser(r6, w)
	expected = "Couldn't find this user."

	if w.String() != expected {
		t.Errorf("PrintUser: The function can't print properly.")
	}
}