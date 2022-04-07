package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, -1},
	}
	for _, testCase := range testCases {
		result := BinarySearch(testCase.nums, testCase.target)
		if result != testCase.want {
			t.Errorf("%#v; want %#v", result, testCase.want)
		}
	}
}
