package array

import (
	"fmt"
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

func TestThreeNumSum(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func TestDailyTemperature(t *testing.T) {
	tem := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(tem))
}

func TestNextArr(t *testing.T) {
	arr := []int{1, 5, 1}
	nextPermutation(arr)
	fmt.Println(arr)
}

func TestSearchRange(t *testing.T) {
	searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
}

func TestMoveZero(t *testing.T) {
	nums := []int{3, 1, 0, 2, 0, 7, 1, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}
