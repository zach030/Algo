package daily

// 统计出现次数
func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := (i + j) / 2
		if nums[m] <= target {
			i = m + 1
			continue
		}
		j = m - 1
	}
	right := i
	if j >= 0 && nums[j] != target {
		return 0
	}
	i, j = 0, len(nums)-1
	for i <= j {
		m := (i + j) / 2
		if nums[m] < target {
			i = m + 1
			continue
		}
		j = m - 1
	}
	left := j
	return right - left - 1
}

// 长度为0-n-1的数组，在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字
// [0,1,3] 长度为2,n=3
func missingNumber(nums []int) int {
	bit := make([]bool, len(nums)+1)
	for i := 0; i < len(bit); i++ {
		bit[i] = false
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= len(bit) {
			continue
		}
		bit[nums[i]] = true
	}
	for i := 0; i < len(bit); i++ {
		if !bit[i] {
			return i
		}
	}
	return -1
}

func missingNumber2(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		n := (i + j) / 2
		if nums[n] == n {
			i = n + 1
		} else {
			j = n - 1
		}
	}
	return i
}

// 二维数组查找,每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序
func findNumberIn2DArray(matrix [][]int, target int) bool {
	x, y := len(matrix)-1, 0
	for x >= 0 && y < len(matrix[0]) {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] < target {
			y++
			continue
		}
		x--
	}
	return false
}

// 旋转数组的最小元素
func minArray(numbers []int) int {

}
