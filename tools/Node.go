package tools

type node[T any] struct {
	data T
	next *node[T]
}

func newNode[T any](data T) *node[T] {
	return &node[T]{data: data}
}

func (n *node[T]) GetData() T {
	return  n.data
}

func (n *node[T]) SetData(number T) {
	n.data = number
}

func (n *node[T]) GetNext() *node[T] {
	return n.next
}

func (n *node[T]) SetNext(next *node[T]) {
	n.next = next
}
