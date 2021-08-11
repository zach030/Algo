package sort

import (
	"fmt"
	"testing"
)

var array = []int{2, 7, 3, 6, 4, 5, 1}

func TestBubbleSort(t *testing.T) {
	ret := BubbleSort(array)
	fmt.Println(ret)
}

func TestSelectSort(t *testing.T) {
	ret := SelectSort(array)
	fmt.Println(ret)
}

func TestInsertSort(t *testing.T) {
	ret := InsertSort(array)
	fmt.Println(ret)
}

func TestQuickSort(t *testing.T) {
	//QuickSort(array, 0, len(array)-1)

	//quickSort(array, 0, len(array)-1)
	sortArray(array)
	fmt.Println(array)
}
