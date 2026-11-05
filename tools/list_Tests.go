package tools

import (
	"testing"
)

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

func TestFindFirst(t *testing.T) {
	list := NewList[int]()

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)

	searcher := NewFindFirstSearcher(NewIntComparable(4))

	if searcher == nil {
		t.Errorf("NewFindFirstSearcher(); Fail creating a FirstSearcher")
	}

	list.Find(searcher)

	if !searcher.Check(4) {
		t.Errorf("Check() = %t; want %t; the method isn't able to compare correctly", true, false)
	}

	if searcher.flag != true {
		t.Errorf("NewFindFirstSearcher(); The searcher didn't find this number: %d", 4)
	}

	if searcher.index != 3 {
		t.Errorf("NewFindFirstSearcher(); The searcher is pointing to a wrong index: %d", searcher.index)
	}
}

func TestFindAll(t *testing.T) {
	list := NewList[int]()

	list.Add(1)
	list.Add(2)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	searcher := NewFindAll(NewIntComparable(2))

	if searcher == nil {
		t.Errorf("NewFindAll(); Fail creating a FindAll")
	}

	list.Find(searcher)

	if !searcher.obj.equals(2) {
		t.Errorf("NewFindAll(); The searcher is searching a wrong number.")
	}

	if searcher.flag != true {
		t.Errorf("NewFindAll(); The searcher didn't find this number: %d", 2)
	}

	if searcher.index[0] != 1 || searcher.index[1] != 3 {
		t.Errorf("NewFindAll(); The searcher is pointing to the wrong index: %d, %d", searcher.index[0], searcher.index[1])
	}
}
