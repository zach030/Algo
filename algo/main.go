package main

import (
	"fmt"
)

func main() {
	var cardNum int64
	fmt.Scan(&cardNum)
	allCards := make([][]int64, 0)
	for i := 0; i < 3; i++ {
		one := make([]int64, 0)
		j := int64(0)
		for j = 0; j < cardNum; j++ {
			var num int64
			fmt.Scan(&num)
			one = append(one, num)
		}
		allCards = append(allCards, one)
	}
	arr1 := SelectSort(allCards[0])
	arr2 := SelectSort(allCards[1])
	arr3 := SelectSort(allCards[2])
	var i int64
	var a int64
	var b int64
	var c int64
	var sum int64
	for i = 0; i < cardNum; i++ {
		arr1,a = getMin(arr1)
		arr2,b = getMin(arr2)
		arr3,c = getMin(arr3)
		sum += getMinest(a, b, c)
	}
	fmt.Println(sum)
}

func getMinest(a, b, c int64) int64 {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if a >= b {
		if b < c {
			return b
		} else {
			return c
		}
	}
	return 0
}

func getMin(arr []int64) ([]int64,int64) {
	ele := arr[0]
	arr = arr[1:]
	return arr,ele
}

func SelectSort(arr []int64) []int64 {
	for i := 0; i < len(arr); i++ {
		pos := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[pos] {
				pos = j
			}
		}
		arr[pos], arr[i] = arr[i], arr[pos]
	}
	return arr
}
