package tree

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 层序遍历
func levelOrder(root *TreeNode) [][]int {
    ret := make([][]int,0)
	// queue用来保存子层级的节点
	queue := make([]*TreeNode,0)
	queue = append(queue, root)
	for len(queue)!=0{
		size := len(queue)
		row := make([]int,0)
		for i:=0;i<size;i++{
			// 取队列头
			node := queue[0]
			row = append(row, node.Val)
			queue = queue[1:]
			if node.Left!=nil{
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