package doubly_linked_list

import "errors"

var ErrorOutOfBounds = errors.New("index is out of bounds")

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
	node := &Node[T]{Value: val, Next: nil, Prev: nil}
	if ll.Length == 0 {
		ll.Head = node
		ll.Tail = node
	} else {
		ll.Tail.Next = node
		node.Prev = ll.Tail
		ll.Tail = node
	}
	ll.Length += 1
}

func (ll *DoublyLinkedList[T]) InsertAt(index int, val T) {
	node := &Node[T]{Value: val, Next: nil, Prev: nil}
	curr := ll.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	curr.Prev.Next = node
	node.Prev = curr.Prev
	node.Next = curr
	curr.Prev = node
	ll.Length += 1
}

func (ll *DoublyLinkedList[T]) Remove(index int) {
	curr := ll.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	curr.Prev.Next = curr.Next
	curr.Next.Prev = curr.Prev
	ll.Length -= 1
}

func (ll *DoublyLinkedList[T]) Get(index int) (T, error) {
	if index >= ll.Length || index < 0 {
		var zero T
		return zero, ErrorOutOfBounds
	}
	curr := ll.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	return curr.Value, nil
}

func (ll *DoublyLinkedList[T]) DeleteNode(node *Node[T]) {
	if node == nil || ll == nil {
		return
	}
	if node == ll.Head {
		ll.Head = node.Next
		ll.Head.Prev = nil
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if node == ll.Tail {
		ll.Tail = node.Prev
		ll.Tail.Next = nil
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	ll.Length--
}
