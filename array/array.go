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

// 三数之和 和为0的组合
func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) < 3 {
		return nil
	}
	for i := 0; i < len(nums); i++ {
		pos := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[pos] {
				pos = j
			}
		}
		nums[pos], nums[i] = nums[i], nums[pos]
	}
	size := len(nums)
	for i := 0; i < size; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, size-1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ret
}
