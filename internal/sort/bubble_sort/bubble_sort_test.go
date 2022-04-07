package sort

import (
	testutil "algorithms/internal/sort/testutil"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	output := testutil.GenerateTestArray(testutil.TestArraySize)
	unsorted := output.Unsorted

	BubbleSort(unsorted)
	for i, v := range output.Sorted {
		if v != unsorted[i] {
			t.Errorf("%#v; want %#v", unsorted[i], v)
		}
	}
}
