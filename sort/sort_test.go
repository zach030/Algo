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
	array := []int{2, 4, 1, 7, 6, 8, 12, 3, 4, 7, 2, 11, 23, 9}
	sortArray(array)
	fmt.Println(array)
}

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func TestDepthOfTree(t *testing.T) {
	tree := &Tree{
		Val: 2,
		Left: &Tree{
			Val: 1,
			Left: &Tree{
				Val: 3,
			},
			Right: nil,
		},
		Right: &Tree{
			Val: 4,
		},
	}
	getDepth(tree)
	//fmt.Println(ret)
}

func getDepth(root *Tree) {
	res := make([]int, 0)
	maxDep := 0
	var dfs func(root *Tree) int
	dfs = func(root *Tree) int {
		if root == nil {
			return 0
		}
		res = append(res, root.Val)
		defer func() { res = res[:len(res)-1] }()
		left := dfs(root.Left)
		right := dfs(root.Right)
		if Max(left, right)+1 > maxDep {
			maxDep = Max(left, right) + 1
			return maxDep
		}
		return maxDep
	}
	dfs(root)
	fmt.Println(res)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type LinkedList struct {
	val  int
	next *LinkedList
}

func TestReverseLinkedList(t *testing.T) {
	head := &LinkedList{val: 1, next: &LinkedList{val: 2, next: &LinkedList{val: 3}}}
	ret := ReverseList(head)
	for ret != nil {
		fmt.Println(ret.val)
		ret = ret.next
	}
}

func ReverseList(head *LinkedList) *LinkedList {
	prev := &LinkedList{}
	curr := head
	for curr != nil {
		tmp := curr.next
		curr.next = prev
		prev = curr
		curr = tmp
	}
	return prev
}

func TestMergeSort(t *testing.T) {
	//n1 := []int{2, 3, 4, 5}
	//n2 := []int{8, 7, 9, 1}
	//arr := merge(n1, n2)
	//fmt.Println(arr)

	nums := []int{2, 4, 3, 5, 8, 7, 9, 1}
	mergeSort(nums, 0, len(nums))
	fmt.Println(nums)

	//MergeSort(nums, 0, len(nums))
	//fmt.Println(nums)
}
