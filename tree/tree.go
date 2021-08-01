package tree

import (
	"fmt"
	"strconv"
)

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
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
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

func calcHeight(root *TreeNode) int {
	if root != nil {
		left := calcHeight(root.Left)
		right := calcHeight(root.Right)
		if left+right > max {
			max = left + right
		}
		return Max(left, right) + 1
	}
	return 0
}

// 路径总和
// DFS
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 叶子节点
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// 路径和
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sumHelp(root, 0)
}

func sumHelp(root *TreeNode, i int) int {
	if root == nil {
		return 0
	}
	tmp := i*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return tmp
	}
	return sumHelp(root.Left, tmp) + sumHelp(root.Right, tmp)
}

func rightSideView(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if i == size-1 {
				ret = append(ret, node.Val)
			}
		}
	}
	return ret
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return Abs(height(root.Left), height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(height(root.Left), height(root.Right)) + 1
}

func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			ret = append(ret, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return ret
}

// 根据前序遍历和中序遍历 构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	index := 0
	for i, val := range inorder {
		if val == rootVal {
			index = i
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:index])+1], inorder[:index])
	root.Right = buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {

	}
	return ret
}

func dfsPathSum(root *TreeNode, sum, target int) {
	if root == nil {
		return
	}
	sum += root.Val

}

// 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 填充next指针，指向右侧节点
func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}
	var connectTwoNode func(node1, node2 *Node)
	connectTwoNode = func(node1, node2 *Node) {
		if node1 == nil || node2 == nil {
			return
		}
		node1.Next = node2
		connectTwoNode(node1.Left, node1.Right)
		connectTwoNode(node1.Right, node2.Left)
		connectTwoNode(node2.Left, node2.Right)
		return
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

// 二叉树展开为链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
	return
}

// 构造最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := 0
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	root := &TreeNode{Val: max}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}

// 寻找重复的子树
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	ret := make([]*TreeNode, 0)
	treeMap := make(map[string]int, 0)

	var traverse func(root *TreeNode) string
	traverse = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		left := traverse(root.Left)
		right := traverse(root.Right)
		subTree := left + "," + right + "," + strconv.Itoa(root.Val)
		fmt.Println("current sub tree is: ", root.Val, ", tree: ", subTree)
		count, ok := treeMap[subTree]
		if ok {
			treeMap[subTree] = count + 1
			if count+1 == 1 {
				ret = append(ret, root)
			}
		} else {
			treeMap[subTree] = 0
		}
		return subTree
	}
	traverse(root)
	return ret
}

// 二叉树中第二小的节点
func findSecondMinimumValue(root *TreeNode) int {
	ans := -1
	val := root.Val
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil || ans != -1 && root.Val >= ans {
			return
		}
		if root.Val > val {
			ans = root.Val
		}
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return ans
}

// 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var judge func(left, right *TreeNode) bool
	judge = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return judge(left.Left, right.Right) && judge(left.Right, right.Left)
	}
	return judge(root.Left, root.Right)
}

// 最大深度
func maxDepth(root *TreeNode) int {
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := depth(root.Left)
		right := depth(root.Right)
		return Max(left, right) + 1
	}
	return depth(root)
}

// 合并二叉树
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1==nil && root2==nil{
		return nil
	}
	if root1==nil{
		return root2
	}
	if root2==nil{
		return root1
	}
	root1.Val+=root2.Val
	root1.Left=mergeTrees(root1.Left,root2.Left)
	root1.Right=mergeTrees(root1.Right,root2.Right)
	return root1
}