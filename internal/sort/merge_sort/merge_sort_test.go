package sort

import (
	testutil "algorithms/internal/sort/testutil"
	"testing"
)

func TestMergeSort(t *testing.T) {
	output := testutil.GenerateTestArray(testutil.TestArraySize)
	unsorted := output.Unsorted

	result := MergeSort(unsorted)

	for i, v := range output.Sorted {
		if v != result[i] {
			t.Errorf("%#v; want %#v", result[i], v)
		}
	}
}
