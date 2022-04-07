package unionfind

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	testCases := []struct {
		size       int
		unionTests []struct {
			x    int
			y    int
			want bool
		}
		findTests []struct {
			x    int
			y    int
			want bool
		}
	}{
		{
			size: 4,
			unionTests: []struct {
				x    int
				y    int
				want bool
			}{
				{0, 1, true},
				{2, 3, true},
				{1, 0, false},
				{3, 2, false},
			},
			findTests: []struct {
				x    int
				y    int
				want bool
			}{
				{0, 1, true},
				{2, 3, true},
				{1, 2, false},
				{0, 2, false},
				{3, 0, false},
				{3, 1, false},
			},
		},
		{
			size: 16,
			unionTests: []struct {
				x    int
				y    int
				want bool
			}{
				{0, 1, true},
				{1, 0, false},
				{1, 2, true},
				{2, 3, true},
				{5, 6, true},
				{7, 1, true},
				{9, 10, true},
				{9, 11, true},
				{9, 13, true},
				{9, 14, true},
				{9, 15, true},
				{1, 15, true},
			},
			findTests: []struct {
				x    int
				y    int
				want bool
			}{
				{0, 1, true},
				{0, 2, true},
				{5, 6, true},
				{7, 0, true},
				{0, 4, false},
				{0, 5, false},
				{4, 5, false},
				{1, 15, true},
			},
		},
	}

	for _, testCase := range testCases {
		uf := New(testCase.size)

		for _, unionTest := range testCase.unionTests {
			result := uf.Union(unionTest.x, unionTest.y)
			if result != unionTest.want {
				t.Errorf("%#v %#v", unionTest.x, unionTest.y)
			}
		}

		for _, findTest := range testCase.findTests {
			result := uf.Find(findTest.x) == uf.Find(findTest.y)
			if result != findTest.want {
				t.Errorf("%#v %#v", findTest.x, findTest.y)
			}
		}
	}
}
