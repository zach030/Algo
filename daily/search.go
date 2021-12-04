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
	//if(j >= 0 && nums[j] != target) return 0;
	//        // 搜索左边界 right
	//        i = 0; j = nums.length - 1;
	//        while(i <= j) {
	//            int m = (i + j) / 2;
	//            if(nums[m] < target) i = m + 1;
	//            else j = m - 1;
	//        }
	//        int left = j;
	//        return right - left - 1;
}
