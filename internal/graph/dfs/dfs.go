package dfs

import (
	"algorithms/internal/graph"
)

// Time complexity: O(V+E)
// Space complexity: O(V+E)
func DFS(g *graph.Graph, src string, dst string) []string {
	path := []string{}
	visitMap := map[string]bool{}
	dfs(g, src, dst, &path, visitMap)
	return path
}

func dfs(g *graph.Graph, node string, dst string, path *[]string, visitMap map[string]bool) bool {
	if visited := visitMap[node]; visited {
		return false
	}

	*path = append(*path, node)
	visitMap[node] = true

	if node == dst {
		return true
	}

	neighbors := g.Neighbors(node)
	for _, neighbor := range neighbors {
		if dfs(g, neighbor, dst, path, visitMap) {
			return true
		}
	}
	return false
}
