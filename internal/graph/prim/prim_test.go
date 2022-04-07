package prim

import (
	"algorithms/internal/graph"
	"reflect"
	"testing"
)

func TestPrim(t *testing.T) {

	g := graph.New(graph.Config{Undirected: true})

	// S: {V: 1, W: 4},
	// V: {S: 1, W: 2, T: 6},
	// W: {T: 3, S: 4, V: 2},
	// T: {V: 6, W: 3},

	g.AddEdge("s", "v", 1)
	g.AddEdge("s", "w", 4)

	g.AddEdge("v", "w", 2)
	g.AddEdge("v", "t", 6)

	g.AddEdge("w", "t", 3)

	mst := Prim(&g)
	result := mst.Connections()
	want := map[string]map[string]int{
		"s": {"v": 1},
		"t": {},
		"v": {"w": 2},
		"w": {"t": 3},
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("%v;\nwant %v", result, want)
	}
}
