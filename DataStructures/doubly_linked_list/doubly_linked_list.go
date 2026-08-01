package doubly_linked_list

type Node[T any] struct {
	Value T
	Next *Node[T]
	Prev *Node[T]
}

type DoublyLinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Length int
}

func New[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		Head: nil,
		Tail: nil,
		Length: 0,
	}
}

func (ll *DoublyLinkedList[T]) Prepend(val T) {
	node := &Node[T]{Value: val, Next: nil, Prev: nil}
	if ll.Length == 0 {
		ll.Head = node
		ll.Tail = node
	} else {
		ll.Head.Prev = node
		node.Next = ll.Head
		ll.Head = node
	}
	ll.Length += 1
}

func (ll *DoublyLinkedList[T]) Append(val T) {
	return
}

func (ll *DoublyLinkedList[T]) InsertAt(index int, val T) {
	return
}

func (ll *DoublyLinkedList[T]) Remove(index int) {
	return
}

func (ll *DoublyLinkedList[T]) Get(index int) (T, error) {
	var zero T
	return zero, nil
}




