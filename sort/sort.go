package sort

import (
	"fmt"
	"math/rand"
	"time"
)

// BubbleSort 冒泡排序
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

func QuickSort(arr []int, left, right int) {
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

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
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
	quickSort(arr, left, i-1)
	quickSort(arr, i+1, right)
}

func sortArray(nums []int) []int {
	//var quickSort func(arr []int,left,right int)
	quickSort := func(arr []int, left, right int) {
		if left >= right {
			return
		}
		pivot := left
		i, j := left, right
		tmp := arr[pivot]
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
		quickSort(arr, left, i-1)
		quickSort(arr, i+1, right)
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	m := (left + right) / 2
	mergeSort(nums, left, m)
	mergeSort(nums, m+1, right)
	nums = merge(nums[left:m], nums[m+1:right])
}

func merge(num1, num2 []int) []int {
	size := len(num1) + len(num2)
	ret := make([]int, size)
	pos := 0
	p, q := 0, 0
	for p < len(num1) && q < len(num2) {
		if num1[p] < num2[q] {
			ret[pos] = num1[p]
			p++
		} else {
			ret[pos] = num2[q]
			q++
		}
		pos++
	}
	if p == len(num1) {
		for i := q; i < len(num2); i++ {
			ret[pos] = num2[i]
			pos++
		}
	}
	if q == len(num2) {
		for i := p; i < len(num1); i++ {
			ret[pos] = num1[i]
			pos++
		}
	}
	fmt.Println("merge in array :", num1, num2)
	fmt.Println("merge out array :", ret)
	return ret
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func MergeSort(arr []int, a, b int) {
	if b-a <= 1 {
		return
	}

	c := (a + b) / 2
	MergeSort(arr, a, c)
	MergeSort(arr, c, b)
	arrLeft := make([]int, c-a)
	arrRight := make([]int, b-c)
	copy(arrLeft, arr[a:c])
	copy(arrRight, arr[c:b])
	i := 0
	j := 0
	for k := a; k < b; k++ {
		if i >= c-a { //超过数组一半
			arr[k] = arrRight[j]
			j++
		} else if j >= b-c {
			arr[k] = arrLeft[i]
			i++
		} else if arrLeft[i] < arrRight[j] {
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j]
			j++
		}
	}
}
