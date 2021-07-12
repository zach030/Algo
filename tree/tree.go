package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历
func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	// queue用来保存子层级的节点
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) != 0 {
		size := len(queue)
		row := make([]int, 0)
		for i := 0; i < size; i++ {
			// 取队列头
			node := queue[0]
			row = append(row, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, row)
	}
	return ret
}

// 锯齿形状层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	// queue用来保存子层级的节点
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	reverse := false
	for len(queue) != 0 {
		size := len(queue)
		row := make([]int, 0)
		for i := 0; i < size; i++ {
			// 取队列头
			node := queue[0]
			row = append(row, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if reverse {
			row = reverseSlice(row)
			ret = append(ret, row)
			reverse = false
		} else {
			reverse = true
			ret = append(ret, row)
		}
	}
	return ret
}

func reverseSlice(s []int) []int {
	size := len(s)
	ret := make([]int, size)
	for i, num := range s {
		ret[size-1-i] = num
	}
	return ret
}

// 最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//todo
	return root
}

var max = 0

// 二叉树直径
func diameterOfBinaryTree(root *TreeNode) int {
	if root != nil {
		_ = calcHeight(root)
		return max
	}
	return 0
}

func calcHeight(root *TreeNode)int {
	if root != nil {
		left := calcHeight(root.Left)
		right := calcHeight(root.Right)
		if left+right>max{
			max = left+right
		}
		return Max(left,right)+1
	}
	return 0
}

func Max(a,b int)int{
	if a >b {
		return a
	}
	return b
}
