package search

// Time complexity: O(log n)
// Space complexity: O(1)
func BinarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := int(uint(l+r) >> 1)
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if nums[l] == target {
		return l
	}

	return -1
}
