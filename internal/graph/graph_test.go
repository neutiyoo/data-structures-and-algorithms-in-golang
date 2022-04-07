package graph

import (
	"reflect"
	"sort"
	"testing"
)

func TestGraphDirected(t *testing.T) {
	g := New(Config{Undirected: false})
	g.AddEdge("a", "b", 1)
	g.AddEdge("b", "c", 2)
	g.AddEdge("b", "d", 3)

	result := g.Connections()
	want := map[string]map[string]int{
		"a": {"b": 1},
		"b": {"c": 2, "d": 3},
		"c": {},
		"d": {},
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("%v;\nwant %v", result, want)
	}

	vetices := g.Vertices()
	for i, v := range []string{"a", "b", "c", "d"} {
		if v != vetices[i] {
			t.Errorf("%#v; want %#v", vetices[i], v)
		}
	}

	neighbors := g.Neighbors("b")
	for i, v := range []string{"c", "d"} {
		if v != neighbors[i] {
			t.Errorf("%#v; want %#v", vetices[i], v)
		}
	}

	w := g.GetEdgeWeight("b", "c")
	if w != 2 {
		t.Errorf("%#v; want %#v", w, 2)
	}

	edges := g.Edges("b")

	for i, v := range []DirectedEdge{
		{X: "b", Y: "c", Weight: 2},
		{X: "b", Y: "d", Weight: 3},
	} {
		if !reflect.DeepEqual(edges[i], v) {
			t.Errorf("%#v; want %#v", edges[i], v)
		}
	}
}

func TestGraphUnirected(t *testing.T) {
	g := New(Config{Undirected: true})
	g.AddEdge("a", "b", 1)
	g.AddEdge("a", "b", 1)
	g.AddEdge("b", "c", 1)
	g.AddEdge("b", "d", 1)

	result := g.Connections()
	want := map[string]map[string]int{
		"a": {"b": 1},
		"b": {"a": 1, "c": 1, "d": 1},
		"c": {"b": 1},
		"d": {"b": 1},
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("%v;\nwant %v", result, want)
	}

	vetices := g.Vertices()
	neighbors := g.Neighbors("b")

	for i, v := range []string{"a", "c", "d"} {
		if v != neighbors[i] {
			t.Errorf("%#v; want %#v", vetices[i], v)
		}
	}
}

func TestSortingDirectedEdges(t *testing.T) {
	g := New(Config{Undirected: true})
	g.AddEdge("a", "a", 1)
	g.AddEdge("a", "b", 3)
	g.AddEdge("a", "c", 2)
	g.AddEdge("a", "d", 4)

	edges := g.Edges("a")

	for i, v := range []int{1, 3, 2, 4} {
		if v != edges[i].Weight {
			t.Errorf("%#v; want %#v", edges[i].Weight, v)
		}
	}

	// Ascending order
	sort.Sort(DirectedEdges(edges))

	for i, v := range []int{1, 2, 3, 4} {
		if v != edges[i].Weight {
			t.Errorf("%#v; want %#v", edges[i].Weight, v)
		}
	}
}
