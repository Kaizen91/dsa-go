package stack

type Node[T any] struct {
	Value T
	Prev *Node[T]
}

type Stack[T any] struct {
	Top *Node[T]
	Length int
}

func (s *Stack[T]) Push(value T) {
	node := &Node[T]{Value: value, Prev: nil}
	if s.Length == 0 {
		s.Top = node
	} else {
		node.Prev = s.Top
		s.Top = node
	}
	s.Length += 1
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.Length == 0 {
		var zero T
		return zero, false
	} else {
		val := s.Top.Value
		s.Top = s.Top.Prev
		s.Length -= 1
		return val, true
	}
}
