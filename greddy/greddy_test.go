package greddy

import (
	"fmt"
	"testing"
)

func Test_findContentChildren(t *testing.T) {
	fmt.Println(findContentChildren([]int{10, 9, 8, 7}, []int{5, 6, 7, 8}))
}

func TestOverlap(t *testing.T) {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}))
}
