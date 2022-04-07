package topologicalsort

import (
	"algorithms/internal/graph"
	"container/list"
)

// Kahn's Algorithm
// Time complexity: O(V+E)
// Space complexity: O(V)
func TopologicalSort(g *graph.Graph) *[]string {
	indegreeMap := map[string]int{}
	vertices := g.Vertices()
	for _, vertex := range vertices {
		indegreeMap[vertex] = 0
	}
	for _, vertex := range vertices {
		for _, neighbor := range g.Neighbors(vertex) {
			indegreeMap[neighbor]++
		}
	}

	queue := list.New()

	// A map is an unordered collection of key-value pairs.
	// Therefore, it is not idempotent.
	for k, v := range indegreeMap {
		if v == 0 {
			queue.PushBack(k)
		}
	}

	order := []string{}

	for queue.Len() > 0 {
		el := queue.Remove(queue.Front())
		vertex := el.(string)
		order = append(order, vertex)
		for _, neighbor := range g.Neighbors(vertex) {
			indegreeMap[neighbor]--
			if indegreeMap[neighbor] == 0 {
				queue.PushBack(neighbor)
			}
		}
	}

	if len(vertices) != len(order) {
		return nil
	}

	return &order
}
