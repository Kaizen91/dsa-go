package graphs
import (
	"github.com/Kaizen91/DSA-Go/DataStructures/queue"
)

type Edge[T comparable] struct {
	u, v T
}

func NewEdge[T comparable](u, v T) *Edge[T] {
	return &Edge[T]{u: u, v: v}
}

type Graph[T comparable] struct {
	start T
	end T
	edges map[T][]Edge[T]
}

func New[T comparable](start, end T) *Graph[T] {
	return &Graph[T]{
		start: start,
		end: end,
		edges: nil,
	}
}

func (g *Graph[T]) BFS() []T {
	q := queue.New[T]()
	seen := make(map[T]bool)
	parent := make(map[T]T)
	found := false
	q.Push(g.start)

	for q.Length != 0 {
		curr, _ := q.Pop()
		if curr == g.end {
			found = true
			break
		}
		for _, edge := range g.edges[curr] {
			if !seen[edge.v] {
				seen[edge.v] = true
				parent[edge.v] = curr
				q.Push(edge.v)
			}
		}
	} 

	if !found {
		return nil
	}

	var path []T
	for curr := g.end; curr != g.start; curr = parent[curr] {
		path = append([]T{curr}, path...)
	}
	path = append([]T{g.start}, path...)
	return path
}

func (g *Graph[T]) walkRecursive(
	curr T,
	seen map[T]bool,
	path *[]T,
) bool {
	*path = append(*path, curr)
	if curr == g.end {
		return true
	seen[curr] = true
	} else {
		for _, edge := range g.edges[curr] {
			if seen[edge.v] {
				return false
			}
			if g.walkRecursive(edge.v, seen, path) {
				return true
			}
		}
	}
	*path = (*path)[:len(*path)-1]
	return false
}

func (g *Graph[T]) DFS() []T {
	seen := make(map[T]bool)
	path := &[]T{}
	g.walkRecursive(g.start, seen, path)
	return *path
}

func (g *Graph[T]) AddEdge(e Edge[T]) {
	if g.edges == nil {
		g.edges = make(map[T][]Edge[T])
	}
	g.edges[e.u] = append(g.edges[e.u], e)
}
