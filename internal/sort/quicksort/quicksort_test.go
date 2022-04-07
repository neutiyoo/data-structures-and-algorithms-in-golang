package sort

import (
	testutil "algorithms/internal/sort/testutil"
	"testing"
)

func TestQuicksort(t *testing.T) {
	output := testutil.GenerateTestArray(testutil.TestArraySize)
	unsorted := output.Unsorted

	Quicksort(&unsorted, 0, len(unsorted)-1)

	for i, v := range output.Sorted {
		if v != unsorted[i] {
			t.Errorf("%#v; want %#v", unsorted[i], v)
		}
	}
}
