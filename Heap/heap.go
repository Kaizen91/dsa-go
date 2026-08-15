package heap

import (
	"container/heap"
	"cmp"
)

type Item[T cmp.Ordered] struct {
	Value T
}

type PriorityQueue[T cmp.Ordered] []*Item[T]

func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Value < pq[j].Value
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(*Item[T])
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue[T]) Init() {
	heap.Init(pq)
}

func (pq *PriorityQueue[T]) PushItem(v T) {
	item := &Item[T]{
		Value: v,
	}
	heap.Push(pq, item)
}

func (pq *PriorityQueue[T]) PopItem() (T, bool) {
	if pq.Len() == 0 {
		var zero T
		return zero, false
	}
	item := heap.Pop(pq).(*Item[T])
	return item.Value, true
}
