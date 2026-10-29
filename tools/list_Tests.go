package tools

import "testing"

func TestInit(t *testing.T) {
	list := NewList[int]()

	if list.GetLength() != 0 {
		t.Errorf("The list didn't initialize empty.")
	}

	if list.rear != nil {
		t.Errorf("The list's rear isn't nil")
	}

	if list.tail != nil {
		t.Errorf("The list's tail isn't nil")
	}
}

func TestLenght(t *testing.T) {
	list := NewList[int]()
	
	if list.GetLength() != 0 {
		t.Errorf("NewList(); The list didn't initialize empty.")
	}

	for i := range 10 {
		list.Add(i)

		if list.GetLength() != (i + 1) {
			t.Errorf("Add(); The list's lenght doesn't match with the expected.")
		}
	}

	for j := range 10 {
		list.Remove(j)

		if list.GetLength() != (99 - j) {
			t.Errorf("Remove(); The list's lenght doesn't match with the expected.")
		}
	}
}
