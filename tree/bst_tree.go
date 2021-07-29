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
