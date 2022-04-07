package bfs

import (
	"algorithms/internal/graph"
	"container/list"
)

// Time complexity: O(V+E)
// Space complexity: O(V)
func BFS(g *graph.Graph, src string, dst string) []string {
	path := []string{}
	visitMap := map[string]bool{}

	queue := list.New()
	queue.PushBack(src)
	for queue.Len() > 0 {
		val := queue.Remove(queue.Front())
		vertex := val.(string)

		path = append(path, vertex)
		visitMap[vertex] = true

		if dst == vertex {
			return path
		}

		for _, neighbor := range g.Neighbors(vertex) {
			if visited := visitMap[neighbor]; !visited {
				queue.PushBack(neighbor)
			}
		}
	}
	return path
}
