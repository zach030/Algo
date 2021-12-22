package daily

import (
	"fmt"
	"testing"
)

func Test_fib(t *testing.T) {
	fmt.Println(fib(5))
}

func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func TestMaxSumSubArr(t *testing.T) {
	a1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(a1))
}

func TestMaxValue(t *testing.T) {
	g := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(maxValue(g))
}
