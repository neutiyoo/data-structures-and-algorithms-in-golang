package sort

// Worst Time complexity: O(n²)
// Average Time complexity: O(n log n)
func Quicksort(nums *[]int, left int, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	pivot := (*nums)[mid]
	index := partition(nums, left, right, pivot)
	Quicksort(nums, left, index-1)
	Quicksort(nums, index, right)
}

func partition(nums *[]int, left int, right int, pivot int) int {
	for left <= right {
		for (*nums)[left] < pivot {
			left++
		}
		for (*nums)[right] > pivot {
			right--
		}
		if left <= right {
			(*nums)[left], (*nums)[right] = (*nums)[right], (*nums)[left]
			left++
			right--
		}
	}
	return left
}
