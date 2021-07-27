package tree

import (
	"fmt"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
			},
		},
	}
	//ret := make([]int,0)
	ret := inorderTraversal(tree)
	fmt.Println(ret)
}
