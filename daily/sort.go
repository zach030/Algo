package daily

import (
	"sort"
	"strconv"
	"strings"
)

type strs []string

func (s strs) Len() int {
	return len(s)
}

func (s strs) Less(i, j int) bool {
	return s[i]+s[j] < s[j]+s[i]
}

func (s strs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 把数组排成最小的数
// 先比第一位，再比第二位，小的放前面
func minNumber(nums []int) string {
	arr := make(strs, len(nums))
	for i, num := range nums {
		arr[i] = strconv.Itoa(num)
	}
	sort.Sort(arr)
	return strings.Join(arr, "")
}

// 判断是否是连续数组
func isStraight(nums []int) bool {
	//todo 12.28
	return true
}

// 最小k个数
func getLeastNumbers(arr []int, k int) []int {
	return nil
}
