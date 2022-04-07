package dijkstra

import (
	"algorithms/internal/graph"
	heap "algorithms/internal/heap"
	"math"
)

const Infinity = math.MaxInt32

// Dijkstra's Algorithm
// Time complexity: O(E log V)
func Dijkstra(g *graph.Graph, sourceVertex string) (map[string]int, map[string]string) {
	dist := make(map[string]int)
	for _, v := range g.Vertices() {
		if v == sourceVertex {
			continue
		}
		dist[v] = Infinity
	}

	prev := make(map[string]string)
	for _, v := range g.Neighbors(sourceVertex) {
		prev[v] = sourceVertex
	}

	mh := heap.Heap{}
	for _, edge := range g.Edges(sourceVertex) {
		dist[edge.Y] = edge.Weight
		mh.Insert(edge.Weight, edge.Y)
	}

	visitMap := make(map[string]bool)

	min := mh.Extract()

	for min != nil {
		shortest := min.Val.(string)
		for _, neighbor := range g.Neighbors(shortest) {
			if neighbor == sourceVertex {
				continue
			}
			alt := dist[shortest] + g.GetEdgeWeight(shortest, neighbor)
			if dist[neighbor] == Infinity || alt < dist[neighbor] {
				dist[neighbor] = alt
				prev[neighbor] = shortest
			}
		}
		visitMap[shortest] = true

		mh = heap.Heap{}
		for k, v := range dist {
			if !visitMap[k] {
				mh.Insert(v, k)
			}
		}
		min = mh.Extract()
	}

	return dist, prev
}
