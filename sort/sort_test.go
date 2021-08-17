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
