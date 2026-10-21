// Ok

package tools

import "errors"

type SLL[T any] struct {
	lenght int
	rear   *node[T]
	tail   *node[T]
}

func(list *SLL[T]) GetRear() *node[T] {
	return list.rear
}

func (sll *SLL[T]) SetRear(node *node[T]) {
	sll.rear = node
}

func(list *SLL[T]) GetTail() *node[T] {
	return list.tail
}

func (sll *SLL[T]) SetTail(node *node[T]) {
	sll.tail = node
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

func (list *SLL[T]) GetLenght() int {
	return list.lenght
}

func (sll *SLL[T]) SetLenght(amount int) {
	sll.lenght = amount
}

func (sll *SLL[T]) Remove(index int) T {
	if sll.GetLenght() <= 0 {
		panic(errors.New("Index out of range"))
	}

	if index < 0 {
		index = index + sll.GetLenght()
	}

	if index == 0 {
		return sll.removeFirst()
	}
	if index == sll.GetLenght() {
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

	sll.SetLenght(sll.GetLenght()-1)

	return data
}

func (sll *SLL[T]) removeLast() T {
	node := sll.GetRear()
	data := sll.GetTail().GetData()

	for range sll.lenght - 1 {
		node = node.GetNext()
	}

	node.SetNext(nil)
	sll.SetTail(node)

	sll.SetLenght(sll.GetLenght()-1)

	return data
}

func (sll *SLL[T]) removeFirst() T {
	number := sll.rear.GetData()
	if sll.GetLenght() == 1 {
		sll.SetRear(nil)
		sll.SetTail(nil)
	} else {
		sll.rear = sll.rear.GetNext()
	}
	sll.SetLenght(sll.GetLenght()-1)
	return number
}

func (sll *SLL[T]) Find(searcher ISearcher[T]) {
	var currentNode = sll.GetRear()
	for index := range sll.GetLenght() {
		searcher.Search(currentNode.GetData(), index)
		if searcher.EarlyStop() {
			return
		}
		currentNode = currentNode.GetNext()
	}
}

func (sll *SLL[T]) SimpleFind(searcher ISimpleSearcher[T]) int {
	var currentNode = sll.GetRear()
	for index := range sll.GetLenght() {
		if searcher.Check(currentNode.GetData()) {
			return index
		}
		currentNode = currentNode.GetNext()
	}

	return -1
}
