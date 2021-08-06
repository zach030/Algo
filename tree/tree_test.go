package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	ret := levelOrder(tree)
	for _, row := range ret {
		fmt.Println(row)
	}
}

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumNumbers(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{Val: 0},
	}
	ret := sumNumbers(root)
	fmt.Println(ret)
}

func TestLowestCommonAncestor(t *testing.T) {
	p := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}
	q := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	root := &TreeNode{
		Val:  3,
		Left: p,
		Right: &TreeNode{
			Val:  1,
			Left: q,
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}
	node := lowestCommonAncestor(root, p, q)

	fmt.Printf("%+v", node)
}

func TestRightSideView(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	ret := rightSideView(tree)
	fmt.Printf("%+v", ret)
}

func TestFindDuplicateTree(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	res := findDuplicateSubtrees(tree)
	fmt.Println(res)
}

func TestFindSecondMiniValue(t *testing.T) {
	tree := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(findSecondMinimumValue(tree))
}

func TestTreeDepth(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   4,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{
				Val:   5,
			},
		},
	}
	fmt.Println(minDepth(tree))
}