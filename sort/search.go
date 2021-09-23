package sort

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
