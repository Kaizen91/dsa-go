package queue

type Node[T any] struct {
	Value T
	Prev *Node[T]
	Next *Node[T]
}

type Queue[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Length int
}

func New[T any]() *Queue[T]{
	return &Queue[T]{
		Head: nil,
		Tail: nil,
		Length: 0,
	}
}

func (q *Queue[T]) Push(value T) {
	node := &Node[T]{Value: value, Prev: nil, Next: nil}
	if q.Length == 0 {
		q.Head = node
		q.Tail = node
	} else {
		node.Next = q.Tail
		q.Tail.Prev = node
		q.Tail = node
	}
	q.Length += 1
}

func (q *Queue[T]) Pop() (T, bool) {
	if q.Length == 0 {
		var zero T
		return zero, false
	} else {
		val := q.Head.Value
		q.Head = q.Head.Prev
		q.Length -= 1
		return val, true
	}
}
