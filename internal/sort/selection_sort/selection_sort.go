package sort

// Time complexity: O(n²)
// Space complexity: O(1)
func SelectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			nums[minIndex], nums[i] = nums[i], nums[minIndex]
		}
	}
}
