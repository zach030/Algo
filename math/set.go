package math

// 递归
// 求子集
func subSet(num []int) [][]int {
	if len(num) == 0 {
		return nil
	}
	last := num[len(num)-1]
	num = num[:len(num)-1]
	prev := subSet(num)
	for i := 0; i < len(prev); i++ {
		prev = append(prev, prev[i])
		prev[len(prev)-1] = append(prev[len(prev)-1], last)
	}
	return prev
}

// 回溯
func subSet2(num []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func(nums []int, start int, track []int)
	backtrack = func(nums []int, start int, track []int) {
		res = append(res, track)
		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack(nums, start+1, track)
			track = track[:len(track)-1]
		}
	}
	backtrack(num, 0, track)
	return res
}
