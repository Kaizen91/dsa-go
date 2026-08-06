package ringbuffer

type RingBuffer[T any] struct {
	Capacity int
	Buffer []T
	Head int
	Tail int
	Length int
}

func New[T any](capacity int) *RingBuffer[T] {
	return &RingBuffer[T]{
		Capacity: capacity,
		Buffer: make([]T, capacity),
		Head: 0,
		Tail: 0,
		Length: 0,
	}
}

func (rb *RingBuffer[T]) Push(item T) error {
	if rb.Length == rb.Capacity {
		return ErrFull
	}
	rb.Buffer[rb.Tail] = item
	rb.Tail = (rb.Tail + 1) % rb.Capacity
	rb.Length += 1
	return nil
}

func (rb *RingBuffer[T]) Pop() (T, error) {
	var zero T
	if rb.Length == 0 {
		return zero, ErrEmpty
	}
	item := rb.Buffer[rb.Head]
	rb.Buffer[rb.Head] = zero
	rb.Head = (rb.Head + 1) % rb.Capacity
	rb.Length -= 1
	return item, nil
}


