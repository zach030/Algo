package tree

import (
	"fmt"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	tree := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}
	//ret := make([]int,0)
	ret := inorderTraversal(tree)
	fmt.Println(ret)

	ret = inorderTraversal2(tree)
	fmt.Println(ret)
}
