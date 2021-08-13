package tree

import "math"

// Binary Search Tree
// 二叉搜索树 对于 BST 的每一个节点node，左子树节点的值都比node的值要小，右子树节点的值都比node的值大

// 寻找第k小
func kthSmallest(root *TreeNode, k int) int {
	// 二叉搜索树中序遍历本身就是升序的
	res := make([]int, 0)
	var inorderTraverse func(root *TreeNode)
	inorderTraverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorderTraverse(root.Left)
		res = append(res, root.Val)
		inorderTraverse(root.Right)
	}
	inorderTraverse(root)
	//fmt.Println(res)
	return res[k-1]
}

// 将二叉搜索树转化为累加树
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var inorderTraverse func(root *TreeNode)
	inorderTraverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorderTraverse(root.Right)
		sum += root.Val
		root.Val = sum
		inorderTraverse(root.Left)
	}
	inorderTraverse(root)
	return root
}

// 不同的二叉搜索树
// 推导出动态方程，动态规划
func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

// 判断有效二叉树
func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

// 二叉树最大路径和
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		innerSum := left + right + root.Val
		maxSum = Max(innerSum, maxSum)
		return Max(root.Val+Max(left, right), 0)
	}
	dfs(root)
	return maxSum
}

// 插入二叉搜索树
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// 删除二叉搜索树节点
func deleteNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		min := getMin(root.Right)
		root.Val = min.Val
		root.Right = deleteNode(root.Right, min.Val)
	} else if root.Val > val {
		root.Left = deleteNode(root.Left, val)
	} else {
		root.Right = deleteNode(root.Right, val)
	}
	return root
}

func getMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// 不同的二叉搜索树
func generateTrees(n int) []*TreeNode {
	ret := make([]*TreeNode, 0)
	if n == 0 {
		return ret
	}
	return build(1, n)
}

func build(low, high int) []*TreeNode {
	ret := make([]*TreeNode, 0)
	if low > high {
		ret = append(ret, nil)
		return ret
	}
	for i := low; i <= high; i++ {
		left := build(low, i-1)
		right := build(i+1, high)
		for _, lt := range left {
			for _, rt := range right {
				root := &TreeNode{Val: i}
				root.Left = lt
				root.Right = rt
				ret = append(ret, root)
			}
		}
	}
	return ret
}

// 最大键值和
func maxSumBST(root *TreeNode) int {
	maxSum := 0
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		var (
			left  int
			right int
			sum   int
		)
		// 判断左右子树是否是二叉搜索树
		if !isValidBST2(*root.Left) || !isValidBST2(*root.Right) {
			goto next
		}
		// 判断加上自身 是否是二叉搜索树
		if root.Val > findMax(*root.Left) || root.Val < findMin(*root.Right) {
			goto next
		}
		left = findSum(*root.Left)
		right = findSum(*root.Right)
		sum = left + right + root.Val
		if sum > maxSum {
			maxSum = sum
		}
		//前序遍历
	next:
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return maxSum
}

// 二叉搜索树最大值
func findMax(root TreeNode) int {
	tree := &root
	for tree.Right != nil {
		tree = tree.Right
	}
	return tree.Val
}

func findMin(root TreeNode) int {
	tree := &root
	for tree.Left != nil {
		tree = tree.Left
	}
	return tree.Val
}

func findSum(root TreeNode) int {
	tree := &root
	sum := 0
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		sum += root.Val
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(tree)
	return sum
}

// 判断有效二叉树
func isValidBST2(root TreeNode) bool {
	val := root
	tree := &val
	stack := make([]*TreeNode, 0)
	inorder := math.MinInt64
	for len(stack) > 0 || tree != nil {
		for tree != nil {
			stack = append(stack, tree)
			tree = tree.Left
		}
		tree = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if tree.Val <= inorder {
			return false
		}
		inorder = tree.Val
		tree = tree.Right
	}
	return true
}
