package tools

import "errors"

type SLL[T any] struct {
	lenght int
	rear   *node[T]
	tail   *node[T]
}

func NewList[T any]() *SLL[T] {
	return &SLL[T]{lenght: 0}
}

func (list *SLL[T]) Add(data T) {
	n := newNode(data)

	if list.lenght == 0 {
		list.rear = n
		list.tail = n
	} else {
		list.tail.next = n
		list.tail = n
	}

	list.lenght++
}

func (list *SLL[T]) Lenght() int {
	return list.lenght
}

func (list *SLL[T]) Remove(index int) (T, bool) {
	lastIndex := list.lenght - 1

	if index > lastIndex || index < 0 || list.rear == nil {
		panic(errors.New("Index error!"))
	}

	if index == 0 {
		deleted := list.rear.data
		list.lenght--

		if list.rear == list.tail {
			list.rear = nil
			list.tail = nil
			return deleted, true
		}

		list.rear = list.rear.next
		return deleted, true
	}

	previousNode := list.rear

	for i := 1; i < index-1; i++ {
		previousNode = previousNode.next
	}

	n := previousNode.next
	deleted := previousNode.next.data

	previousNode.next = n.next
	list.lenght--

	return deleted, true
}

/* 
func (list *SLL[T]) Find(data any) T {
	previousNode := list.rear
	lastIndex := list.lenght - 1

	for i := 0; i < lastIndex; i++ {
		listData := previousNode.data
		previousNode = previousNode.next

		if data == listData {
			return i

		}
	}

	return -1
}
*/