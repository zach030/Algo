package tree

// 中序遍历
// 1. 递归
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root==nil{
			return
		}
		inorder(root.Left)
		ret = append(ret,root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return ret
}

// 2. 迭代
func inorderTraversal2(root *TreeNode)[]int{
	ret := make([]int,0)

	return ret
}