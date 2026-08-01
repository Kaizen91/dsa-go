package data_structures

type Node[T any] struct {
	Value T
	Next *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Length int
}

func (ll *LinkedList[T]) Push(val T) {
	node := Node[T]{Value: val, Next: ll.Head}
	ll.Head = &node
	ll.Length += 1
}

func (ll *LinkedList[T]) Pop() (T, bool) {
	if ll.Length == 0 {
		var zero T
		return zero, false
	}
	node := *ll.Head
	ll.Head = node.Next
	ll.Length -= 1
	if ll.Head == nil {
		ll.Tail = nil
	}
	return node.Value, true
}


