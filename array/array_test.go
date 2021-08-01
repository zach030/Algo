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
