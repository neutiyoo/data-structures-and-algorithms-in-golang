package sort

// Time complexity: O(n log n)
// Space complexity: O(1)
func HeapSort(nums []int) {
	// Build max heap
	for i := len(nums) / 2; i > -1; i-- {
		maxHeapifyDown(nums, i, len(nums))
	}

	for i := len(nums) - 1; i > -1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapifyDown(nums, 0, i)
	}
}

func maxHeapifyDown(nums []int, parent int, length int) {
	larger := parent
	left := parent*2 + 1
	right := parent*2 + 2

	if left < length && nums[left] > nums[larger] {
		larger = left
	}
	if right < length && nums[right] > nums[larger] {
		larger = right
	}
	if larger != parent {
		nums[parent], nums[larger] = nums[larger], nums[parent]
		maxHeapifyDown(nums, larger, length)
	}
}
