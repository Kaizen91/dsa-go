package graphs

import (
	"testing"
	"reflect"
)


type testCase[T comparable] struct {
	name string
	graph *Graph[T]
	expectedBFS []T
	expectedDFS []T
}

func GenerateSharedTestCases() []testCase[int] {
	graph := New(0,5)
	graph.AddEdge(Edge[int]{u: 0, v: 1})
	graph.AddEdge(Edge[int]{u: 0, v: 4})
	graph.AddEdge(Edge[int]{u: 1, v: 4})
	graph.AddEdge(Edge[int]{u: 1, v: 2})
	graph.AddEdge(Edge[int]{u: 2, v: 5})
	graph.AddEdge(Edge[int]{u: 3, v: 5})
	graph.AddEdge(Edge[int]{u: 5, v: 4})
	return []testCase[int]{
		{
			name: "basic graph",
			graph: graph,
			expectedBFS: []int{0, 1, 2, 5},
			expectedDFS: []int{0, 1, 2, 3, 5},
		},
	}
}

func TestBFS(t *testing.T) {
	tests := GenerateSharedTestCases()
	for _, test := range tests {
		t.Run(test.name + "_BFS", func(t *testing.T) {
			got := test.graph.BFS()
			if !reflect.DeepEqual(got, test.expectedBFS) {
				t.Errorf("BFS() = %v want %v", got, test.expectedBFS)
			}
		})
	}
}

func TestDFS(t *testing.T) {
	tests := GenerateSharedTestCases()
	for _, test := range tests {
		t.Run(test.name + "_DFS", func(t *testing.T) {
			got := test.graph.DFS()
			if !reflect.DeepEqual(got, test.expectedDFS) {
				t.Errorf("DFS() = %v want %v", got, test.expectedDFS)
			}
		})
	}
}

