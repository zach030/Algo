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
	if root==nil || root==p || root == q{
		return root
	}
	left := lowestCommonAncestor(root.Left,p,q)
	right := lowestCommonAncestor(root.Right,p,q)
	if left == nil  && right == nil {
		return nil
	}
	if left==nil{
		return right
	}
	if right==nil{
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



// 路径总和
// DFS
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root==nil{
		return false
	}
	// 叶子节点
	if root.Left==nil && root.Right == nil {
		return root.Val==targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}


// 路径和
func sumNumbers(root *TreeNode) int {
	if root==nil{
		return 0
	}
	return sumHelp(root, 0)
}

func sumHelp(root *TreeNode, i int)int{
	if root==nil{
        return 0
    }
	tmp := i*10 + root.Val
	if root.Left==nil && root.Right==nil{
		return tmp
	}
	return sumHelp(root.Left, tmp)+sumHelp(root.Right, tmp) 
}

func rightSideView(root *TreeNode) []int {
	ret := make([]int,0)
	if root==nil{
		return ret
	}
	queue := make([]*TreeNode,0)
	queue = append(queue,root)
	for len(queue)!=0{
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left!=nil{
				queue = append(queue,node.Left)
			}
			if node.Right!=nil{
				queue = append(queue,node.Right)
			}
			if i==size-1{
				ret = append(ret,node.Val)
			}
		}
	}
	return ret
}

func isBalanced(root *TreeNode) bool {
	if root==nil{
		return true
	}
	return Abs(height(root.Left),height(root.Right))<=1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode)int{
	if root==nil{
		return 0
	}
	return Max(height(root.Left),height(root.Right))+1
}

func preorderTraversal(root *TreeNode) []int {
	ret := make([]int,0)
	if root==nil{
		return ret
	}
	stack := make([]*TreeNode,0)
	for root!=nil || len(stack)!=0{
		for root!=nil{
			ret = append(ret,root.Val)
			stack = append(stack,root)
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
	if len(preorder)==0{
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	index := 0
	for i, val := range inorder {
		if val==rootVal{
			index = i
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:index])+1],inorder[:index])
	root.Right = buildTree(preorder[len(inorder[:index])+1:],inorder[index+1:])
	return root
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ret := make([][]int,0)
	if root == nil{
		return nil
	}
	if root.Left==nil && root.Right==nil{

	}
	return ret
}

func dfsPathSum(root *TreeNode, sum,target int){
	if root==nil{
		return
	}
	sum += root.Val

}