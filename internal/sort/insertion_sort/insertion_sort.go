package sort

// Time complexity: O(n²)
// Space complexity: O(1)
func InsertionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		j := i
		for j > 0 && nums[j-1] > temp {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = temp
	}
}
