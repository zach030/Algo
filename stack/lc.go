package stack

// 字符串解码
func decodeString(s string) string {
	//for i := 0; i < len(s); i++ {
	//	ch := s[i]
	//
	//}
	return ""
}

var result [][]int

func backtrack(nums, pathNums []int, used []bool) {
	if len(nums) == len(pathNums) {
		tmp := make([]int, len(nums))
		copy(tmp, pathNums)
		result = append(result, tmp)
		//result=append(result,pathNums)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			pathNums = append(pathNums, nums[i])
			backtrack(nums, pathNums, used)
			pathNums = pathNums[:len(pathNums)-1]
			used[i] = false
		}
	}
}

// 496
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	greater := nextGreater(nums2)
	ret := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				ret[i] = greater[j]
			}
		}
	}
	return ret
}

func nextGreater(nums []int) []int {
	res := make([]int, len(nums))
	stack := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return res
}

// 环形数组
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	stack := make([]int, 0)
	for i := 2*n - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums[i%n] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i%n] = -1
		} else {
			res[i%n] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i%n])
	}
	return res
}
