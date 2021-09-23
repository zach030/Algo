package sort

import (
	"fmt"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	arr := []int{1, 4, 6, 7, 9, 12}
	target := 6
	fmt.Println(binarySearch(arr, target))

	fmt.Println(bsSearchInternal(arr, 0, len(arr)-1, target))
}
