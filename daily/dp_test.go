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
