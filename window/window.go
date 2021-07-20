package window

import "fmt"

// 无重复字符最长子串
func lengthOfLongestSubstring(s string) int {
	size := len(s)
	max := 0
	start, end := 0, 0
	filter := make(map[string]int, 0)
	for end < size {
		if nextIndex, ok := filter[string(s[end])]; ok {
			if nextIndex > start {
				start = nextIndex
			}
		}
		if end-start+1 > max {
			max = end - start + 1
		}
		filter[string(s[end])] = end + 1
		end++
	}
	return max
}

// 最短无序连续子数组

func findUnsortedSubarray(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	left := 0
	right := 0
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < max {
			left = i
			break
		}
		max = nums[i]
	}
	min := nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > min {
			right = i
			break
		}
		min = nums[i]
	}
	fmt.Println(left, right)
	if left == right {
		return 0
	}
	return right - left + 3
}

func findUnsortedSubarray2(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	size := len(nums)
	high, low := 0, size-1
	max, min := nums[0], nums[size-1]
	for i := 1; i < size; i++ {
		max = Max(max, nums[i])
		min = Min(min, nums[size-1-i])
		if nums[i] < max {
			high = i
		}
		if nums[size-1-i] > min {
			low = size - 1 - i
		}
	}
	if high > low {
		return high - low + 1
	}
	return 0
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
