package topologicalsort

import (
	"algorithms/internal/graph"
	"reflect"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	testCases := []struct {
		edges          [][]string
		possibleOrders [][]string
	}{
		{
			// DAG
			[][]string{
				{"0", "1"},
			},
			[][]string{
				{"0", "1"},
			},
		},
		{
			// DAG
			[][]string{
				{"0", "1"},
				{"0", "2"},
				{"1", "3"},
				{"2", "3"},
			},
			[][]string{
				{"0", "1", "2", "3"},
				{"0", "2", "1", "3"},
			},
		},
		{
			// Not a DAG
			[][]string{
				{"0", "1"},
				{"0", "2"},
				{"1", "3"},
				{"2", "3"},
				{"3", "0"},
			},
			nil,
		},
		{
			// DAG
			[][]string{
				{"2", "3"},
				{"3", "1"},
				{"4", "0"},
				{"4", "1"},
				{"5", "0"},
				{"5", "2"},
			},
			[][]string{
				{"4", "5", "0", "2", "3", "1"},
				{"5", "4", "2", "0", "3", "1"},
			},
		},
	}

	for _, testCase := range testCases {
		g := graph.New(graph.Config{Undirected: false})
		for _, edge := range testCase.edges {
			g.AddEdge(edge[0], edge[1], 1)
		}
		result := TopologicalSort(&g)
		if result != nil {
			isValid := false
			for _, order := range testCase.possibleOrders {
				if reflect.DeepEqual(*result, order) {
					isValid = true
				}
			}
			if !isValid {
				t.Errorf("got %#v", result)
			}
		} else {
			if !reflect.DeepEqual(result, (*[]string)(nil)) {
				t.Errorf("%#v; want %#v", result, (*[]string)(nil))
			}
		}
	}
}
