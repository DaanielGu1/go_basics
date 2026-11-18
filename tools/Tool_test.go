package tools

import (
	"Treino-em-go/models"
	"testing"
)

func TestInit(t *testing.T) {
	list := NewList[int]()

	if list.GetLength() != 0 {
		t.Errorf("SLL.GetLength() = %d; want %d", list.GetLength(), 0)
	}

	if list.rear != nil {
		t.Errorf("SLL.rear is not initialized as nil (axiom error)")
	}

	if list.tail != nil {
		t.Errorf("SLL.tail is not initialized as nil (axiom error)")
	}

	firstSearcher := NewFindFirstSearcher(NewStringComparable("sample"))

	if firstSearcher.flag != false {
		t.Errorf("firstSearcher.flag = %t; want %t; flag should be false as default", false, true)
	}

	if firstSearcher.index != 0 {
		t.Errorf("firstSearcher.flag = %d; want %d; index should be 0 as default", firstSearcher.index, 0)
	}

	if !firstSearcher.obj.equals("sample") {
		t.Errorf(`firstSearcher.obj = %d; want "sample"; the method isn't able to compare correctly`, firstSearcher.obj)
	}

	allSearcher := NewFindAll(NewStringComparable("sample"))

	if allSearcher.flag != false {
		t.Errorf("allSearcher.flag = %t; want %t; flag should be false as default", false, true)
	}

	if allSearcher.index != nil {
		t.Errorf("allSearcher.flag = %d; want nil; index should be nil as default", allSearcher.index)
	}

	if !allSearcher.obj.equals("sample") {
		t.Errorf(`allSearcher.obj = %d; want "sample"; the method isn't able to compare correctly`, allSearcher.obj)
	}

}

func TestLength(t *testing.T) {
	list := NewList[int]()

	if list.GetLength() != 0 {
		t.Errorf("SLL.GetLength() = %d; want %d", list.GetLength(), 0)
	}

	for i := range 100 {
		list.Add(i)

		if list.GetLength() != (i + 1) {
			t.Errorf("SLL.GetLength() = %d; want %d; the method Add is not correctly implemented; call number error: %d", list.GetLength(), i, i)
		}
	}

	for i := range 100 {
		list.Remove(0)

		if list.GetLength() != (99 - i) {
			t.Errorf("SLL.GetLength() = %d; want %d; the method Remove is not correctly implemented; call number error: %d", list.GetLength(), 99-i, i)
		}
	}
}

func TestAddAndRemove(t *testing.T) {
	list := NewList[int]()

	list.Add(1)

	if list.rear != list.tail {
		t.Errorf("SLL.Add(); first call of method should set rear and tail as same node, but they aren't.")
	}

	list.Add(2)
	if list.rear == list.tail {
		t.Errorf("SLL.Add(); second call of method should set rear and tail as different nodes, but they aren't.")
	}

	if list.rear.next != list.tail {
		t.Errorf("SLL.Add(); second call of method should set rear.next and tail as same nodes, but they aren't.")
	}

	list.Add(3)
	// [1, 2, 3]

	if list.rear == list.tail {
		t.Errorf("SLL.Add(); third call of method should set rear and tail as different nodes, but they aren't.")
	}

	if list.rear.next.next != list.tail {
		t.Errorf("SLL.Add(); third call of method should set rear.next.next and tail as same nodes, but they aren't.")
	}

	if list.length != 3 {
		t.Errorf("SLL.Add() = %d; want 3; the length doesn't matches with the expected.", list.length)
	}

	obj := list.Remove(0)
	if obj != 1 {
		t.Errorf("SLL.Remove() = %d; want %d; the method got issues while trying to remove the first element", obj, 5)
	}

	if list.length != 2 {
		t.Errorf("SLL.Remove() = %d; want 2; the length doesn't matches with the expected.", list.length)
	}

	list.Add(obj)
	// [2, 3, 1]

	obj = list.Remove(-2)
	if obj != 2 {
		t.Errorf("SLL.Remove() = %d; want %d; the method got issues while trying to remove last element", obj, 5)
	}

	list.Add(obj)
	// [2, 1, 3]

	obj = list.Remove(1)
	// [2, 3]

	if obj != 1 {
		t.Errorf("SLL.Remove() = %d; want %d; the method got issues while trying to remove the middle element", obj, 5)
	}

	if list.tail == nil {
		t.Errorf("SLL.Remove(): the method didn't set tail after remove.")
	}
}

func TestGet(t *testing.T) {
	list := NewList[int]()

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)

	obj := list.Get(1)

	if obj != 2 {
		t.Errorf("SLL.Get() = %d; want %d; the method got issues while trying to get the expected element", obj, 2)
	}

	obj = list.Get(-2)

	if obj != 5 {
		t.Errorf("SLL.Get() = %d; want %d; the method got issues while trying to get the expected element", obj, 2)
	}
}

func TestAllSearcher(t *testing.T) {
	list := NewList[string]()

	list.Add("")
	list.Add("123")
	list.Add("hello")
	list.Add("sample")
	list.Add("this")
	list.Add("sample")
	list.Add("456")
	list.Add("789")

	allSearcher := NewFindAll(NewStringComparable("sample"))

	list.Find(allSearcher)
	if allSearcher.index[0] != 3 || allSearcher.index[1] != 5 {
		t.Errorf(`Find(); allSearcher.index[0] = %d; allSearcher.index[1] = %d; want %d and %d; the method 
		got issues while trying to get the expected objects index`, allSearcher.index[0], allSearcher.index[1], 3, 5)
	}
	if len(allSearcher.index) != 2 {
		t.Errorf("Find(); len(allSearcher.index) = %d; want %d; the method got issues while trying to get the expected amount of objects", len(allSearcher.index), 2)
	}

	index := list.SimpleFind(allSearcher)

	if index != 3 {
		t.Errorf("SimpleFind() = %d; want %d; the method got issues while trying to get the expected object index", index, 3)
	}

	if index == -1 {
		t.Errorf("SimpleFind() = -1; want 3: the method didn't found the item in the list.")
	}
}

func TestFirstSearcher(t *testing.T) {
	list := NewList[string]()

	list.Add("")
	list.Add("123")
	list.Add("hello")
	list.Add("sample")
	list.Add("this")
	list.Add("sample")
	list.Add("456")
	list.Add("789")

	firstSearcher := NewFindFirstSearcher(NewStringComparable("sample"))

	if !firstSearcher.Check("sample") {
		t.Errorf("Check() = %t; want %t; the method isn't able to compare correctly", true, false)
	}

	list.Find(firstSearcher)
	if firstSearcher.index != 3 {
		t.Errorf("Find(); firstSearcher.index = %d; want %d;", firstSearcher.index, 3)
	}

	index := list.SimpleFind(firstSearcher)
	if index != 3 {
		t.Errorf("SimpleFind() = %d; want %d; the method got issues while trying to get the expected object index", index, 3)
	}

	firstSearcher.Search("sample", 1)

	if firstSearcher.index != 1 {
		t.Errorf("firstSearcher.index = %d; want %d; the method isn't capable of comparing correctly", firstSearcher.index, 3)
	}

}

func TestWalk(t *testing.T) {
	sll := NewList[int]()
	sll.Add(3)
	sll.Add(4)
	sll.Add(1)
	sll.Add(2)

	rear := sll.rear

	shouldBeNil := rear.Walk(-1)

	if shouldBeNil != nil {
		t.Errorf("Walk() = %p; want nil; the method shouldn't return any value", shouldBeNil)
	}

	shouldBeNil = rear.Walk(100)

	if shouldBeNil != nil {
		t.Errorf("Walk() = %p; want nil; the method shouldn't return any value", shouldBeNil)
	}

	rear = rear.Walk(0)

	if rear.data != 3 {
		t.Errorf("Walk() = %d; want %d; the method got issues while trying to get to the expected node when wanted node == 0", rear.data, 3)
	}

	node := rear.Walk(2)

	if node.data != 1 {
		t.Errorf("Walk() = %d; want %d; the method got issues while trying to get to the expected node when wanted node != 0", node.data, 1)
	}

	number, flag := node.WalkGetData(1)

	if !flag {
		t.Errorf("WalkGetData() = (%d , %t); want (%d , %t); the method failed to get the data", number, flag, node.data, true)
	}

	number, flag = node.WalkGetData(100)
	if flag {
		t.Errorf("WalkGetData() = (%d , %t); want (%d , %t); the method shouldn't return any value", number, flag, 2, false)
	}

	if !node.WalkSetNext(1) {
		t.Errorf("WalkSetNext() = %t; want %t; the method should set node's next attribute to the next node", false, true)
	}

	if node.WalkSetNext(7) {
		t.Errorf("WalkSetNext() = %t; want %t; the method shouldn't be able to set node's next attribute to any value", node.WalkSetNext(7), false)
	}

}

func TestNewComparable(t *testing.T) {
	if NewIntComparable(5) == nil {
		t.Errorf("NewIntComparable(): the comparable is not being created.")
	}

	if NewStringComparable("sample") == nil {
		t.Errorf("NewStringComparable(): the comparable is not being created.")
	}

	group := models.CreateNewGroup(0, "sample")

	if NewGroupComparable(group) == nil {
		t.Errorf("NewGroupComparable(): the comparable is not being created.")
	}

	groupList := []models.Group{*group}

	if NewUserComparable(models.CreateNewUser(0, "JhonDoe", groupList, "123")) == nil {
		t.Errorf("NewStringComparable(): the comparable is not being created.")
	}
}

func TestEquals(t *testing.T) {
	group := models.CreateNewGroup(0, "sample")
	groupComparable := NewGroupComparable(group)
	groupList := []models.Group{*group}
	user := models.CreateNewUser(0, "JhonDoe", groupList, "123")
	userComparable := NewUserComparable(models.CreateNewUser(0, "JhonDoe", groupList, "123"))
	intComparable := NewIntComparable(5)
	stringComparable := NewStringComparable("sample")
	findFirstComparable := NewFindFirstSearcher(intComparable)

	if !groupComparable.equals(group) {
		t.Errorf("equals(): the comparable didn't verify the same group of the parameter.")
	}

	if !userComparable.equals(user) {
		t.Errorf("equals(): the comparable didn't verify the same user of the parameter.")
	}

	if !intComparable.equals(5) {
		t.Errorf("equals(): the comparable didn't verify the same int of the parameter.")
	}

	if !stringComparable.equals("sample") {
		t.Errorf("equals(): the comparable didn't verify the same string of the parameter.")
	}

	if !findFirstComparable.equals(5) {
		t.Errorf("equals(): the comparable didn't verify the same thing of the parameter.")
	}
}