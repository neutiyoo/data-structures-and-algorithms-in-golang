package graph

import (
	"sort"
)

type DirectedEdge struct {
	X      string
	Y      string
	Weight int
}

type DirectedEdges []DirectedEdge

func (d DirectedEdges) Len() int {
	return len(d)
}
func (d DirectedEdges) Swap(x, y int) {
	d[x], d[y] = d[y], d[x]
}
func (d DirectedEdges) Less(x, y int) bool {
	return d[x].Weight < d[y].Weight
}

type Config struct {
	Undirected bool
}
type Graph struct {
	connections map[string]map[string]int
	undirected  bool
	vetices     []string
	neighbors   map[string][]string
	edges       map[string][]DirectedEdge
}

func New(config Config) Graph {
	g := Graph{}
	g.undirected = config.Undirected
	g.connections = make(map[string]map[string]int)
	g.vetices = []string{}
	g.neighbors = make(map[string][]string)
	g.edges = make(map[string][]DirectedEdge)
	return g
}

func (g *Graph) addVertex(vertex string) {
	if g.connections[vertex] == nil {
		g.connections[vertex] = make(map[string]int)
		g.setVertices()
	}
}

func (g *Graph) AddEdge(x string, y string, weight int) {
	if g.connections[x] == nil {
		g.addVertex(x)
	}
	if g.connections[y] == nil {
		g.addVertex(y)
	}

	g.connections[x][y] = weight
	g.setNeighbors(x)
	g.setEdges(x)

	if g.undirected {
		g.connections[y][x] = weight
		g.setNeighbors(y)
		g.setEdges(y)
	}

}
func (g *Graph) setVertices() {
	keys := make([]string, 0, len(g.connections))
	for k := range g.connections {
		keys = append(keys, k)
	}
	/*
		When iterating over a map with a range loop,
		the iteration order is not specified and
		is not guaranteed to be the same from one iteration to the next.
	*/
	sort.Strings(keys)
	g.vetices = keys
}

func (g *Graph) setNeighbors(vertex string) {
	keys := make([]string, 0, len(g.connections[vertex]))
	for k := range g.connections[vertex] {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	g.neighbors[vertex] = keys
}

func (g *Graph) setEdges(vertex string) {
	neighbors := g.Neighbors(vertex)
	edges := make([]DirectedEdge, 0, len(neighbors))
	for _, neighbor := range neighbors {
		edges = append(edges, DirectedEdge{
			X:      vertex,
			Y:      neighbor,
			Weight: g.GetEdgeWeight(vertex, neighbor),
		})
	}
	g.edges[vertex] = edges
}

func (g *Graph) Connections() map[string]map[string]int {
	return g.connections
}

func (g *Graph) Vertices() []string {
	return g.vetices
}

func (g *Graph) Neighbors(vertex string) []string {
	return g.neighbors[vertex]
}

func (g *Graph) GetEdgeWeight(x string, y string) int {
	return g.connections[x][y]
}

func (g *Graph) Edges(vertex string) []DirectedEdge {
	return g.edges[vertex]
}
