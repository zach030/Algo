package tree

import (
	"fmt"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	tree := &TreeNode{
		Val:   3,
		Left:  &TreeNode{
			Val:   1,
			Right: &TreeNode{
				Val:   2,
			},
		},
		Right: &TreeNode{
			Val:   4,
		},
	}
	fmt.Println(kthSmallest(tree,1))
}

func TestConvertTree(t *testing.T) {
	tree := &TreeNode{
		Val:   4,
		Left:  &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 2,
				Right: &TreeNode{Val:   3}},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 7,
				Right: &TreeNode{Val: 8}},
		},
	}
	n := convertBST(tree)
	fmt.Printf("%+v",n)
}