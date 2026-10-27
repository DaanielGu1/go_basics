// Ok

package tools

import "errors"

type SLL[T any] struct {
	length int
	rear   *node[T]
	tail   *node[T]
}

func (list *SLL[T]) GetRear() *node[T] {
	return list.rear
}

func (sll *SLL[T]) SetRear(node *node[T]) {
	sll.rear = node
}

func (list *SLL[T]) GetTail() *node[T] {
	return list.tail
}

func (sll *SLL[T]) SetTail(node *node[T]) {
	sll.tail = node
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
	node := sll.GetRear()

	for range index - 1 {
		node = node.GetNext()
	}

	data := node.GetNext().GetData()

	node.SetNext(node.GetNext().GetNext())

	sll.SetLength(sll.GetLength() - 1)

	return data
}

func (sll *SLL[T]) removeLast() T {
	node := sll.GetRear()
	data := sll.GetTail().GetData()

	for range sll.length - 1 {
		node = node.GetNext()
	}

	node.SetNext(nil)
	sll.SetTail(node)

	sll.SetLength(sll.GetLength() - 1)

	return data
}

func (sll *SLL[T]) removeFirst() T {
	number := sll.rear.GetData()
	if sll.GetLength() == 1 {
		sll.SetRear(nil)
		sll.SetTail(nil)
	} else {
		sll.rear = sll.rear.GetNext()
	}
	sll.SetLength(sll.GetLength() - 1)
	return number
}

func (sll *SLL[T]) Find(searcher ISearcher[T]) {
	var currentNode = sll.GetRear()
	for index := range sll.GetLength() {
		searcher.Search(currentNode.GetData(), index)
		if searcher.EarlyStop() {
			return
		}
		currentNode = currentNode.GetNext()
	}
}

func (sll *SLL[T]) SimpleFind(searcher ISimpleSearcher[T]) int {
	var currentNode = sll.GetRear()
	for index := range sll.GetLength() {
		if searcher.Check(currentNode.GetData()) {
			return index
		}
		currentNode = currentNode.GetNext()
	}

	return -1
}

func (sll *SLL[T]) GetElementByIndex(index int) *node[T] {
	currentNode := sll.GetRear()
	
	for range index {
		currentNode = currentNode.GetNext()
	}

	return currentNode
}
