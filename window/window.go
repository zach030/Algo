package window

import (
	"fmt"
	"math"
)

// 无重复字符最长子串
func lengthOfLongestSubstring(s string) int {
	size := len(s)
	start, end := 0, 0
	max := 0
	filter := make(map[string]int, 0)
	for end < size {
		if nextIndex, ok := filter[string(s[end])]; ok {
			start = nextIndex
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

// 最小覆盖子串
func minWindow(s string, t string) string {
	need, window := make(map[uint8]int, 0), make(map[uint8]int, 0)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	valid := 0
	left, right := 0, 0
	start, size := 0, math.MaxInt32
	for right < len(s) {
		char := s[right]
		right++
		if _, ok := need[char]; ok {
			window[char]++
			if window[char] == need[char] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left < size {
				start = left
				size = right - left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if size == math.MaxInt32 {
		return ""
	}
	return s[start : start+size]
}

// 是否包含字符串排序
func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[uint8]int, 0), make(map[uint8]int, 0)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		char := s2[right]
		right++
		if _, ok := need[char]; ok {
			window[char]++
			if window[char] == need[char] {
				valid++
			}
		}
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			ch := s2[left]
			left++
			if _, ok := need[ch]; ok {
				if window[ch] == need[ch] {
					valid--
				}
				window[ch]--
			}
		}
	}
	return false
}

func findAnagrams(s string, p string) []int {
	need, window := make(map[uint8]int), make(map[uint8]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	left, right := 0, 0
	valid := 0
	ret := make([]int, 0)
	for right < len(s) {
		char := s[right]
		right++
		if _, ok := need[char]; ok {
			window[char]++
			if window[char] == need[char] {
				valid++
			}
		}
		for right-left >= len(p) {
			if valid == len(need) {
				ret = append(ret, left)
			}
			ch := s[left]
			left++
			if _, ok := need[ch]; ok {
				if window[ch] == need[ch] {
					valid--
				}
				window[ch]--
			}
		}
	}
	return ret
}

// 无重复最长字串
func lengthOfLongestSubstring2(s string) int {
	maxSize := 0
	left, right := 0, 0
	window := make(map[uint8]int)
	for right < len(s) {
		ch := s[right]
		right++
		window[ch]++
		for window[ch] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		maxSize = Max(maxSize,right-left)
	}
	return maxSize
}
