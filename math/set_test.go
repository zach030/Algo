package math

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_subSet(t *testing.T) {
	//fmt.Println(subSet([]int{1, 2, 3}))
	//fmt.Println(subSet2([]int{1, 2, 3}))

	fmt.Println(largestNumber([]int{999999998, 999999997, 999999999}))

}

func largestNumber(nums []int) string {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if strconv.Itoa(nums[i])+strconv.Itoa(nums[j]) < strconv.Itoa(nums[j])+strconv.Itoa(nums[i]) {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	fmt.Println(nums)
	str := ""
	for _, i := range nums {
		str = fmt.Sprintf("%s%v", str, i)
	}
	if str[0] == '0' {
		return "0"
	}
	return str
}
