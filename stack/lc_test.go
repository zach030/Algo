package stack

import (
	"fmt"
	"testing"
)

func Test_nextGreater(t *testing.T) {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreater([]int{2, 1, 2, 4, 3}))
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}
