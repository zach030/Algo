package search

// 有序数组中查找元素的第一个和最后一个位置
// 第一个大于等于target的元素下标，第一个大于元素的下标-1
func searchRange(nums []int, target int) []int {
	left := binarySearch(nums, target, true)
	right := binarySearch(nums, target, false) - 1
	if left <= right && right < len(nums) && nums[left] == target && nums[right] == target {
		return []int{left, right}
	}
	return []int{-1, -1}
}

func binarySearch(nums []int, target int, eq bool) int {
	left, right := 0, len(nums)-1
	var res = len(nums)
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target || (eq && nums[mid] >= target) {
			right = mid - 1
			res = mid
		} else {
			left = mid + 1
		}
	}
	return res
}

// 旋转数组查找target
// [nums = [4,5,6,7,0,1,2], target = 0]
// 将数组一分为二，一个一定有序，另一个可能无序，但切分后其中一个必定有序
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			// 如果0--mid有序，且target在有序范围内，则right=mid-1
			//               如果right不在有序范围内，则left=mid+1
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 如果mid--right有序，且target在有序范围内，则left=mid+1
			//                   如果target不在有序范围内，则right=mid-1
			if nums[mid] < target && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// 矩阵搜索，每一行都有序，每行的开头都比前一行的末尾元素大
func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	left, right := 0, row*col-1
	for left <= right {
		mid := (left + right) / 2
		x, y := pos(row, col, mid)
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// (3,4) 8 --(2,0)
func pos(row, col, idx int) (int, int) {
	return idx / col, idx % col
}

// 旋转数组找最小值
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		pivot := (right-left)/2 + left
		if nums[pivot] < nums[right] {
			right = pivot
		}
		if nums[pivot] > nums[right] {
			left = pivot + 1
		}
	}
	return nums[left]
}

// 寻找峰值，严格大于左右两个元素
func findPeakElement(nums []int) int {
	return 0
}
