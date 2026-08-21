package lru

import (
	"sync"
	"github.com/Kaizen91/DSA-Go/DataStructures/doubly_linked_list"
)

type entry[K comparable, V any] struct {
	key K
	value V
}

type LruCache[K comparable, V any] struct {
	mu sync.Mutex
	capacity int
	items map[K]*doubly_linked_list.Node[entry[K, V]]
	evictList *doubly_linked_list.DoublyLinkedList[entry[K, V]]
}

func New[K comparable, V any](capacity int) *LruCache[K, V] {
	return &LruCache[K, V]{
		capacity: capacity,
		items: make(map[K]*doubly_linked_list.Node[entry[K, V]]),
		evictList: doubly_linked_list.New[entry[K, V]](),
	}
}

func (lru *LruCache[K, V]) Get(key K) (V, bool) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	if elem, exists := lru.items[key]; exists {
		lru.evictList.DeleteNode(elem)
		lru.evictList.Prepend(elem.Value)
		lru.items[key] = lru.evictList.Head
		return elem.Value.value, true
	}
	var zero V
	return zero, false
}

func (lru *LruCache[K, V]) Put(key K, val V) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	entry := entry[K, V]{ key: key, value: val}
	if elem, exists := lru.items[key]; exists {
		lru.evictList.DeleteNode(elem)
		lru.evictList.Prepend(entry)
		lru.items[key] = lru.evictList.Head
	} else {
		lru.evictList.Prepend(entry)
		lru.items[key] = lru.evictList.Head
	}

	if lru.capacity < lru.evictList.Length {
		lru.trimCache()
	}

}

func (lru *LruCache[K, V]) trimCache() {
	elem := lru.evictList.Tail
	if elem == nil {
		return
	}
	delete(lru.items, elem.Value.key)
	lru.evictList.DeleteNode(elem)
}
