package array

// 寻找第K个最大元素
func findKthLargest(nums []int, k int) int {
	// 选择排序，每次选出最小的，选k次
	min := 0
	for i := 0; i < len(nums)-k+1; i++ {
		pos := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[pos] {
				pos = j
			}
		}
		nums[pos], nums[i] = nums[i], nums[pos]
		min = nums[i]
	}
	return min
}
