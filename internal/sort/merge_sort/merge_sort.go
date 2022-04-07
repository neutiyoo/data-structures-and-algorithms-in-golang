package sort

// Time complexity: O(n log n)
// Space complexity: O(n)
func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	subX := nums[:mid]
	subY := nums[mid:]
	return merge(MergeSort(subX), MergeSort(subY))
}

func merge(x []int, y []int) []int {
	sorted := []int{}
	for len(x) != 0 && len(y) != 0 {
		if x[0] < y[0] {
			sorted = append(sorted, x[0])
			x = x[1:]
		} else {
			sorted = append(sorted, y[0])
			y = y[1:]
		}
	}
	sorted = append(sorted, x...)
	sorted = append(sorted, y...)
	return sorted
}
