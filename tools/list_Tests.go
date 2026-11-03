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

func TestAddAndRemove(t *testing.T) {
	list := NewList[int]()

	list.Add(0)

	if list.rear != list.tail {
		t.Errorf("Add(); The list's rear and tail doesn't match in the first call of the method.")
	}

	list.Add(1)

	if list.rear == list.tail {
		t.Errorf("Add(); The list's rear and tail still match in the second call of the method.")
	}

	if list.rear.next != list.tail {
		t.Errorf("Add(); The list's rear.next doesn't match with tail in the second call of the method.")
	}

	list.Add(2)

	if list.rear == list.tail {
		t.Errorf("Add(); The list's rear and tail still match in the third call of the method.")
	}

	if list.rear.next.next != list.tail {
		t.Errorf("Add(); The list's rear.next.next doesn't match with tail in the third call of the method.")
	}

	obj := list.Remove(0)

	if obj != 0 {
		t.Errorf("Remove(); The removed item (%d) doesn't match with the first item (%d).", obj, 0)
	}

	list.Add(obj)
	
	obj = list.Remove(2)
	
	if obj != 0 {
		t.Errorf("Remove(); The removed item (%d) doesn't match with the last item (%d).", obj, 0)
	}

	list.Add(3)

	obj = list.Remove(1)
	
	if obj != 2 {
		t.Errorf("Remove(); The removed item (%d) doesn't match with the middle item (%d).", obj, 2)
	}
}

func TestSearcher(t *testing.T) {
	list := NewList[string]()

	list.Add("")
	list.Add("123")
	list.Add("hello")
	list.Add("sample")
	list.Add("this")
	list.Add("sample")
}
	
