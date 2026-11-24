package in

import (
	"testing" 
	"strings"
)

func TestGetString(t *testing.T) {
	var r = strings.NewReader("Daniel")
	var w = &strings.Builder{}

	name, err := IOGetString("message", r, w)

	if err != nil {
		t.Error(err)
	}

	if name != "Daniel" {
		t.Errorf("IOGetString() = %s; want Daniel: The function got the a wrong string.", name)
	}

	if name == "" {
		t.Errorf("IoGetString(): The function got an empty string.")
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
		t.Errorf("IOGetString() = %s; want Daniel: The function got the a wrong string.", num)
	}

	if num == 0 {
		t.Errorf("IoGetString(): The function got an empty string.")
	}
}