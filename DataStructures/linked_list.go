package data_structures

type Node[T any] struct {
	Value T
	Next *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
	Length int
}

func (ll *LinkedList[T]) Push(val T) int {
	return 1
}

func (ll *LinkedList[T]) Pop() (int, bool) {
	return 1, true
}


