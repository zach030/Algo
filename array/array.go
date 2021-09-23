package array

import (
	"math/rand"
	"sort"
	"time"
)

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

// 每日气温
func dailyTemperatures(temperatures []int) []int {
	size := len(temperatures)
	ret := make([]int, size)
	ret[size-1] = 0
	for i := size - 2; i >= 0; i-- {
		for j := i + 1; j < size; j += ret[j] {
			if temperatures[i] < temperatures[j] {
				ret[i] = j - i
				break
			} else if ret[j] == 0 {
				ret[i] = 0
				break
			}
		}
	}
	return ret
}

// 寻找第K个最大元素
func findKthLargest2(nums []int, k int) int {
	rand.Seed(time.Now().Unix())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, index int) int {
	q := randomPartition(nums, left, right)
	if q == index {
		return nums[q]
	} else if q > index {
		return randomPartition(nums, q+1, right)
	}
	return randomPartition(nums, left, q-1)
}

func randomPartition(nums []int, left, right int) int {
	r := rand.Intn(right-left) + left
	nums[left], nums[r] = nums[r], nums[left]
	return partition(nums, left, right)
}

func partition(nums []int, left, right int) int {
	x := nums[right]
	i := left - 1
	for j := left; j < right; j++ {
		if nums[j] <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}

// 下一个排列
func nextPermutation(nums []int) {
	front := search(nums)
	if front >= 0 {
		back := searchBigger(nums, front)
		nums[front], nums[back] = nums[back], nums[front]
	}
	left, right := front+1, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func search(arr []int) int {
	size := len(arr)
	front := size - 2
	for front >= 0 && arr[front] >= arr[front+1] {
		front--
	}
	return front
}

func searchBigger(arr []int, target int) int {
	size := len(arr)
	front := size - 1
	for front > 0 && arr[front] <= arr[target] {
		front--
	}
	return front
}

// 奇数位于偶数前面
func exchange(nums []int) []int {
	f, l := 0, len(nums)-1
	for f < l {
		if nums[f]%2 != 0 {
			f++
			continue
		}
		if nums[l]%2 == 0 {
			l--
			continue
		}
		nums[f], nums[l] = nums[l], nums[f]
		f++
		l--
	}
	return nums
}

// 只出现一次
func singleNumbers(nums []int) []int {
	record := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		if _, ok := record[nums[i]]; ok {
			delete(record, nums[i])
			continue
		}
		record[nums[i]] = true
	}
	ret := make([]int, 0)
	for val, _ := range record {
		ret = append(ret, val)
	}
	return ret
}

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for nums[left]+nums[right] != target {
		if nums[left]+nums[right] > target {
			right--
			continue
		}
		if nums[left]+nums[right] < target {
			left++
			continue
		}
		break
	}
	return []int{nums[left], nums[right]}
}

// 拼接最小的数
func minNumber(nums []int) string {
	return ""
}

func findContinuousSequence(target int) [][]int {
	left, right, half := 1, 2, target/2+1
	ret := make([][]int, 0)
	for left <= half {
		if calcSum(left, right) < target {
			right++
			continue
		}
		if calcSum(left, right) > target {
			left++
			continue
		}
		arr := make([]int, 0, right-left)
		for i := left; i <= right; i++ {
			arr = append(arr, i)
		}
		ret = append(ret, arr)
		left++
	}
	return ret
}

func calcSum(a, b int) int {
	return (a + b) * (b - a + 1) / 2
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	area := 0
	left, right := 0, len(height)-1
	for left < right {
		if height[left] <= height[right] {
			area = max(area, height[left]*(right-left))
			left++
		} else {
			area = max(area, height[right]*(right-left))
			right--
		}
	}
	return area
}

func searchRange(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for j < len(nums) {
		nums[j] = 0
		j++
	}
}
