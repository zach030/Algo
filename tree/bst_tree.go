package tree

// Binary Search Tree
// 二叉搜索树 对于 BST 的每一个节点node，左子树节点的值都比node的值要小，右子树节点的值都比node的值大

// 寻找第k小
func kthSmallest(root *TreeNode, k int) int {
	// 二叉搜索树中序遍历本身就是升序的
	res := make([]int,0)
	var inorderTraverse func(root *TreeNode)
	inorderTraverse = func(root *TreeNode) {
		if root==nil{
			return
		}
		inorderTraverse(root.Left)
		res = append(res,root.Val)
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
		if root==nil{
			return
		}
		inorderTraverse(root.Right)
		sum+=root.Val
		root.Val =sum
		inorderTraverse(root.Left)
	}
	inorderTraverse(root)
	return root
}