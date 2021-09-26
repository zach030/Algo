package sort

import "fmt"

func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			low = mid + 1
			continue
		}
		if nums[mid] > target {
			high = mid - 1
		}
	}
	return -1
}

func bsSearchInternal(nums []int, low, high, value int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	if nums[mid] == value {
		return mid
	}
	if nums[mid] < value {
		return bsSearchInternal(nums, mid+1, high, value)
	}
	if nums[mid] > value {
		return bsSearchInternal(nums, low, mid+1, value)
	}
	return -1
}

func findKMax(nums []int, left, right, k int) {
	pos := partition(nums, left, right)
	fmt.Println("pos is:", pos, ", k is:", k, " ,arr is:", nums)
	if pos == k-1 {
		fmt.Println(nums[pos])
	} else if pos > k-1 {
		findKMax(nums, left, pos-1, k)
	} else {
		findKMax(nums, pos+1, right, k-pos)
	}
}

func partition(nums []int, left, right int) int {
	tmp := nums[left]
	for left < right {
		for left < right && tmp >= nums[right] {
			right--
		}
		nums[left] = nums[right]
		for left < right && tmp <= nums[left] {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = tmp
	return right
}
