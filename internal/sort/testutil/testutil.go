package testutil

type TestArray struct {
	Sorted   []int
	Unsorted []int
}

const TestArraySize = 1024

func GenerateTestArray(size int) TestArray {
	sorted := make([]int, size)
	unsorted := []int{}
	m := make(map[int]*int)
	for i := range sorted {
		sorted[i] = i
		m[i] = nil
	}
	for k := range m {
		unsorted = append(unsorted, k)
	}
	return TestArray{Sorted: sorted, Unsorted: unsorted}
}
