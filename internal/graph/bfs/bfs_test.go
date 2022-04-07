package bfs

import (
	"algorithms/internal/graph"
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	g := graph.New(graph.Config{Undirected: false})

	g.AddEdge("a", "b", 1)
	g.AddEdge("b", "a", 1)
	g.AddEdge("b", "c", 1)
	g.AddEdge("b", "d", 1)
	g.AddEdge("c", "b", 1)
	g.AddEdge("c", "f", 1)
	g.AddEdge("d", "b", 1)
	g.AddEdge("d", "e", 1)
	g.AddEdge("e", "d", 1)
	g.AddEdge("f", "c", 1)

	testCases := []struct {
		src  string
		dst  string
		want []string
	}{
		{"a", "z", []string{"a", "b", "c", "d", "f", "e"}},
		{"a", "f", []string{"a", "b", "c", "d", "f"}},
	}
	for _, testCase := range testCases {
		result := BFS(&g, testCase.src, testCase.dst)
		if !reflect.DeepEqual(result, testCase.want) {
			t.Errorf("%#v; want %#v", result, testCase.want)
		}
	}
}
