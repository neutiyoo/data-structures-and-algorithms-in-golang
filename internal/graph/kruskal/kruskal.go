package kruskal

import (
	"algorithms/internal/graph"
	unionfind "algorithms/internal/union_find"
	"sort"
)

// Kruskal's Algorithm
// Time complexity: O(E log V)
func Kruskal(g *graph.Graph) *graph.Graph {

	mst := graph.New(graph.Config{Undirected: false})

	allEdges := make([]graph.DirectedEdge, 0)
	for _, v := range g.Vertices() {
		allEdges = append(allEdges, g.Edges(v)...)
	}

	// Ascending order
	sort.Sort(graph.DirectedEdges(allEdges))

	indexMap := make(map[string]int)
	for i, v := range g.Vertices() {
		indexMap[v] = i
	}

	forest := unionfind.New(len(indexMap))

	for _, v := range allEdges {
		xIndex := indexMap[v.X]
		yIndex := indexMap[v.Y]
		if forest.Find(xIndex) != forest.Find(yIndex) {
			forest.Union(xIndex, yIndex)
			mst.AddEdge(v.X, v.Y, v.Weight)
		}
	}

	return &mst
}
