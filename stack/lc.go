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
