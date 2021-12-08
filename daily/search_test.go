package daily

import (
	"fmt"
	"testing"
)

func Test_missingNumber(t *testing.T) {
	arr1 := []int{0, 1, 3}
	fmt.Println(missingNumber(arr1))
	arr2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println(missingNumber(arr2))
}

func TestFindInMatrix(t *testing.T) {
	matrix := [][]int{{1, 2}, {3, 4}}
	fmt.Println(matrix[1][0])
	fmt.Println(findNumberIn2DArray(matrix, 2))
}

func TestMinArr(t *testing.T) {
	a := minArray([]int{3, 4, 5, 1, 2})
	fmt.Println(a)
}

func TestUniqCha(t *testing.T) {
	fmt.Println(string(firstUniqChar("loveleetcode")))
}

func TestLevelOrder(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(levelOrder(tree))
}
