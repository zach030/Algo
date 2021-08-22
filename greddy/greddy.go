package greddy

import "sort"

// 贪心算法
// 455 g胃口，s尺寸（s>g）
func findContentChildren(g []int, s []int) int {
	// 先排序
	sort.Ints(g)
	sort.Ints(s)
	child, cookie := 0, 0
	for child < len(g) && cookie < len(s) {
		if g[child] <= s[cookie] {
			child++
		}
		cookie++
	}
	return child
}

// 135 分发糖果
func candy(ratings []int) int {
	// 先初始化每个人分1个
	arr := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		arr[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			arr[i] = arr[i-1] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			arr[i-1] = max(arr[i-1], arr[i]+1)
		}
	}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 435 移除最少区间，保证不重叠
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	total := 0
	prev := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < prev {
			total++
		} else {
			prev = intervals[i][1]
		}
	}
	return total
}

// 605 种花问题
func canPlaceFlowers(flowerbed []int, n int) bool {

	return true
}