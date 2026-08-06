package ringbuffer

import (
	"testing"
	"errors"
	"fmt"
)

var (
	ErrEmpty = errors.New("ringbuffer: buffer is empty")
	ErrFull = errors.New("ringbuffer: buffer is full")
)

type step[T any] struct {
	op string
	val T
	wantVal T
	wantErr error
	expHead int
	expTail int
	expLength int
}

func assertState[T any](t *testing.T, rb *RingBuffer[T], expHead, expTail, expLength int) {
	t.Helper()
	if rb.Head != expHead || rb.Tail != expTail || rb.Length != expLength {
		t.Errorf("State mismatch!\n Got: head=%d, tail=%d, length=%d\n" +
				 "Want: head=%d, tail=%d, length=%d\n",
				 rb.Head,
				 rb.Tail,
				 rb.Length,
				 expHead,
				 expTail,
				 expLength,
		)
	}
}

func TestRingBuffer_StepSequence(t *testing.T) {
	rb := New[int](3)
	steps := []step[int]{
		//Pop on empty buffer
		{op: "pop", wantErr: ErrEmpty, expHead: 0, expTail: 0, expLength: 0},
		//Fill to capacity
		{op: "push", val: 10, expHead: 0, expTail: 1, expLength: 1},
		{op: "push", val: 20, expHead: 0, expTail: 2, expLength: 2},
		{op: "push", val: 30 , expHead: 0, expTail: 0, expLength: 3},
		//Push to full buffer
		{op: "push", val: 40, wantErr: ErrFull, expHead: 0, expTail: 0, expLength: 3},
		//Drain buffer
		{op: "pop", wantVal: 10, expHead: 1, expTail: 0, expLength: 2},
		{op: "pop", wantVal: 20, expHead: 2, expTail: 0, expLength: 1},
		{op: "pop", wantVal: 30, expHead: 0, expTail: 0, expLength: 0},
		//Pop on empty buffer
		{op: "pop", wantErr: ErrEmpty, expHead: 0, expTail: 0, expLength: 0},
	}

	for i, s := range steps {
		t.Run(fmt.Sprintf("Step_%d_%s,", i, s.op), func(t *testing.T) {
			if s.op == "push" {
				err := rb.Push(s.val)
				if !errors.Is(err, s.wantErr) {
					t.Fatalf("Push(%d) err = %v, want = %v", s.val, err, s.wantErr)
				}
			} else if s.op == "pop" {
				got, err := rb.Pop()
				if !errors.Is(err, s.wantErr) {
					t.Fatalf("Pop() err = %v, want = %v", err, s.wantErr)
				}
				if s.wantErr == nil && got != s.wantVal {
					t.Errorf("Pop() = %v, want = %v", got, s.wantVal)
				}
			}
			assertState(t, rb, s.expHead, s.expTail, s.expLength)
		})
	}
}
