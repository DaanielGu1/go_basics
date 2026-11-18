// Ok

package tools

import "errors"

type SLL[T any] struct {
	length int
	rear   *node[T]
	tail   *node[T]
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
		panic(errors.New("Empty list"))
	}

	if index < 0 {
		index = sll.GetLength() + index

		if index < 0 {
			panic(errors.New("Index out of range"))
		}
	} else if index > sll.GetLength()-1 {
		panic(errors.New("Index out of range"))
	}

	sll.SetLength(sll.GetLength() - 1)

	if index == 0 {
		data := sll.rear.GetData()
		sll.rear = sll.rear.GetNext()

		if sll.length == 0 {
			sll.tail = nil
		}

		return data
	}

	node := sll.rear.Walk(index - 1)
	data, ok := node.WalkGetData(1)
	if !ok {
		panic(errors.New("Index out of range"))
	}
	node.WalkSetNext(1)

	if index == sll.GetLength() {
		//VALOR DA LISTA ANTERIORMENTE FOI REDUZIDO EM 1!
		sll.tail = node
	}

	return data
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
