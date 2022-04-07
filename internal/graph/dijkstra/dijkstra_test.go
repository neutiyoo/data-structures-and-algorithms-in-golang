package dijkstra

import (
	"algorithms/internal/graph"
	"testing"
)

func TestDijkstra(t *testing.T) {
	g := graph.New(graph.Config{Undirected: false})

	//     S: {V: 1, W: 4},
	//     V: {W: 2, T: 6},
	//     W: {T: 3},
	//     T: {},

	g.AddEdge("s", "v", 1)
	g.AddEdge("s", "w", 4)

	g.AddEdge("v", "s", 1)
	g.AddEdge("v", "w", 2)
	g.AddEdge("v", "t", 6)

	g.AddEdge("w", "t", 3)

	dist, prev := Dijkstra(&g, "s")

	testCases := []struct {
		src      string
		distWant int
		prevWant string
	}{
		{"t", 6, "w"},
		{"v", 1, "s"},
		{"w", 3, "v"},
	}
	for _, testCase := range testCases {
		distResult := dist[testCase.src]
		if distResult != testCase.distWant {
			t.Errorf("%#v; want %#v", distResult, testCase.distWant)
		}

		prevResult := prev[testCase.src]
		if prevResult != testCase.prevWant {
			t.Errorf("%#v; want %#v", prevResult, testCase.prevWant)
		}
	}
}
