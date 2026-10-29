// Ok

package tools

import "errors"

type SLL[T any] struct {
	length int
	rear   *node[T]
	tail   *node[T]
}

func (list *SLL[T]) setRear() *node[T] {
	return list.rear
}

func (sll *SLL[T]) getRear(data T) {
	sll.rear.data = data
}

func (list *SLL[T]) getTail() *node[T] {
	return list.tail
}

func (sll *SLL[T]) setTail(data T) {
	sll.tail.data = data
}

func NewList[T any]() *SLL[T] {
	return &SLL[T]{length: 0}
}

func (list *SLL[T]) Add(data T) {
	n := newNode(data)

	if list.length == 0 {
		list.rear = n
		list.tail = n
	} else {
		list.tail.next = n
		list.tail = n
	}

	list.length++
}

func (list *SLL[T]) GetLength() int {
	return list.length
}

func (sll *SLL[T]) SetLength(amount int) {
	sll.length = amount
}

func (sll *SLL[T]) Remove(index int) T {
	if sll.GetLength() <= 0 {
		panic(errors.New("Index out of range"))
	}

	if index < 0 {
		index = index + sll.GetLength()
	}

	if index == 0 {
		return sll.removeFirst()
	}
	if index == sll.GetLength() {
		return sll.removeLast()
	}
	return sll.removeMiddle(index)
}

func (sll *SLL[T]) removeMiddle(index int) T {
	node := sll.rear

	for range index - 1 {
		node = node.GetNext()
	}

	data := node.GetNext().GetData()

	node.SetNext(node.GetNext().GetNext())

	sll.SetLength(sll.GetLength() - 1)

	return data
}

func (sll *SLL[T]) removeLast() T {
	current := sll.rear
	data := sll.tail.data

	for range sll.length - 1 {
		current = current.GetNext()
	}
	
	sll.tail = current
	current.SetNext(nil)

	sll.SetLength(sll.GetLength() - 1)

	return data
}

func (sll *SLL[T]) removeFirst() T {
	number := sll.rear.GetData()
	if sll.GetLength() == 1 {
		sll.rear = nil
		sll.tail = nil
	} else {
		sll.rear = sll.rear.GetNext()
	}
	sll.SetLength(sll.GetLength() - 1)
	return number
}

func (sll *SLL[T]) Find(searcher ISearcher[T]) {
	var currentNode = sll.rear
	for index := range sll.GetLength() {
		searcher.Search(currentNode.GetData(), index)
		if searcher.EarlyStop() {
			return
		}
		currentNode = currentNode.GetNext()
	}
}

func (sll *SLL[T]) SimpleFind(searcher ISimpleSearcher[T]) int {
	var currentNode = sll.rear
	for index := range sll.GetLength() {
		if searcher.Check(currentNode.GetData()) {
			return index
		}
		currentNode = currentNode.GetNext()
	}

	return -1
}

func (sll *SLL[T]) Get(index int) T {
	if sll.GetLength() <= 0 {
		panic(errors.New("Empty list"))
	}

	if index < 0 {
		index = sll.GetLength() + index

		if index < 0 {
			panic(errors.New("Index out of range"))
		} else if index > sll.GetLength() - 1 {
			panic(errors.New("Index out of range"))
		}
	}

	node := sll.rear
	for range index {
		node = node.next
	}

	return node.data
}

