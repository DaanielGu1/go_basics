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

func (n *node[T]) GetNext() *node[T] {
	return n.next
}

func (n *node[T]) SetNext(next *node[T]) {
	n.next = next
}

func (n *node[T]) Walk(steps int) *node[T] {
	if steps < 0 {
		return nil
	}
	if steps == 0 {
		return n
	}

	node := n
	for range steps {
		node = node.GetNext()
		if node == nil {
			return nil
		}
	}

	return node
}

func (n *node[T]) WalkGetData(steps int) (T, bool) {
	node := n.Walk(steps)
	if node != nil {
		return node.GetData(), true
	}

	return n.GetData(), false
}

func (n *node[T]) WalkSetNext(steps int) bool {
	node := n.Walk(steps)
	if node == nil {
		return false
	}

	n.SetNext(node.GetNext())
	return true
}
