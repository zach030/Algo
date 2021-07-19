package sort

import (
	"math/rand"
	"time"
)

// 冒泡排序
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		exchange := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				exchange = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !exchange {
			return arr
		}
	}
	return arr
}

// 选择排序
func SelectSort(arr []int) []int {
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

// 插入排序
func InsertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		cur := arr[i]
		j := i - 1
		for j >= 0 && cur < arr[j] {
			arr[j+1], arr[j] = arr[j], arr[j+1]
			j--
		}
	}
	return arr
}

// 希尔排序
func ShellSort(arr []int) []int {

	return arr
}

func QuickSort(arr []int, left, right int){
	if left >= right {
		return
	}
	i := left
	j := right
	rand.Seed(time.Now().Unix())
	r := rand.Intn(right-left) + left
	arr[i], arr[r] = arr[r], arr[i]
	tmp := arr[i]
	for i < j {
		for i < j && arr[j] >= tmp {
			j--
		}
		arr[i] = arr[j]
		for i < j && arr[i] <= tmp {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = tmp
	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}
