package prim

import (
	"algorithms/internal/graph"
	"algorithms/internal/heap"
)

// Prim's Algorithm
// Time complexity: O(E log V)
func Prim(g *graph.Graph) *graph.Graph {
	mst := graph.New(graph.Config{Undirected: false})

	// Choose arbitrarily
	src := g.Vertices()[0]
	visitMap := map[string]bool{
		src: true,
	}

	mh := heap.Heap{}
	for _, v := range g.Edges(src) {
		mh.Insert(v.Weight, v)
	}

	for mh.Len() > 0 {
		min := mh.Extract()
		cur := min.Val.(graph.DirectedEdge)

		var nextVertex string
		if !visitMap[cur.X] {
			nextVertex = cur.X
		}
		if !visitMap[cur.Y] {
			nextVertex = cur.Y
		}

		if len(nextVertex) != 0 {
			mst.AddEdge(cur.X, cur.Y, cur.Weight)
			visitMap[nextVertex] = true

			for _, v := range g.Edges(nextVertex) {
				if !visitMap[v.X] || !visitMap[v.Y] {
					mh.Insert(v.Weight, v)
				}
			}
		}

	}

	return &mst
}
