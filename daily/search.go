package daily

// 统计出现次数
func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := (i + j) / 2
		if nums[m] <= target {
			i = m + 1
			continue
		}
		j = m - 1
	}
	right := i
	if j >= 0 && nums[j] != target {
		return 0
	}
	i, j = 0, len(nums)-1
	for i <= j {
		m := (i + j) / 2
		if nums[m] < target {
			i = m + 1
			continue
		}
		j = m - 1
	}
	left := j
	return right - left - 1
}

// 长度为0-n-1的数组，在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字
// [0,1,3] 长度为2,n=3
func missingNumber(nums []int) int {
	bit := make([]bool, len(nums)+1)
	for i := 0; i < len(bit); i++ {
		bit[i] = false
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= len(bit) {
			continue
		}
		bit[nums[i]] = true
	}
	for i := 0; i < len(bit); i++ {
		if !bit[i] {
			return i
		}
	}
	return -1
}

func missingNumber2(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		n := (i + j) / 2
		if nums[n] == n {
			i = n + 1
		} else {
			j = n - 1
		}
	}
	return i
}

// 二维数组查找,每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序
func findNumberIn2DArray(matrix [][]int, target int) bool {
	x, y := len(matrix)-1, 0
	for x >= 0 && y < len(matrix[0]) {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] < target {
			y++
			continue
		}
		x--
	}
	return false
}

// 旋转数组的最小元素
// [3,4,5,1,2]
func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		m := (left + right) / 2
		if numbers[m] > numbers[right] {
			left = m + 1
			continue
		}
		if numbers[m] < numbers[right] {
			right = m
			continue
		}
		right--
	}
	return numbers[left]
}

//输入：s = "abaccdeff"
//输出：'b'
func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return byte(' ')
	}
	bm := make(map[uint8]bool)
	for i := 0; i < len(s); i++ {
		if _, ok := bm[s[i]]; ok {
			bm[s[i]] = false
			continue
		}
		bm[s[i]] = true
	}
	for i := 0; i < len(s); i++ {
		if bm[s[i]] {
			return s[i]
		}
	}
	return ' '
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	ret := make([]int, 0)
	for len(nodes) != 0 {
		head := nodes[0]
		ret = append(ret, head.Val)
		nodes = nodes[1:]
		if head.Left != nil {
			nodes = append(nodes, head.Left)
		}
		if head.Right != nil {
			nodes = append(nodes, head.Right)
		}
	}
	return ret
}

//从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。每一层打一行
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	ret := make([][]int, 0)
	for len(nodes) != 0 {
		size := len(nodes)
		row := make([]int, 0)
		for i := 0; i < size; i++ {
			head := nodes[0]
			nodes = nodes[1:]
			row = append(row, head.Val)
			if head.Left != nil {
				nodes = append(nodes, head.Left)
			}
			if head.Right != nil {
				nodes = append(nodes, head.Right)
			}
		}
		ret = append(ret, row)
	}
	return ret
}

// 层序遍历，第一行从左到右，第二行从右到左
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	ret := make([][]int, 0)
	flag := false
	for len(nodes) != 0 {
		size := len(nodes)
		row := make([]int, 0)
		for i := 0; i < size; i++ {
			head := nodes[0]
			nodes = nodes[1:]
			row = append(row, head.Val)
			if head.Left != nil {
				nodes = append(nodes, head.Left)
			}
			if head.Right != nil {
				nodes = append(nodes, head.Right)
			}
		}
		if flag {
			row = reverse(row)
			flag = false
		} else {
			flag = true
		}
		ret = append(ret, row)
	}
	return ret
}

func reverse(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}

// 树的子结构
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func recur(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}

func mirrorTree(root *TreeNode) *TreeNode {
	var fun func(root *TreeNode)
	fun = func(root *TreeNode) {
		if root == nil {
			return
		}
		left := root.Left
		root.Left = root.Right
		root.Right = left
		fun(root.Left)
		fun(root.Right)
	}
	fun(root)
	return root
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(left, right *TreeNode) bool
	dfs = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if right.Val != left.Val {
			return false
		}
		return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}
	return dfs(root.Left, root.Right)
}

// 矩阵board中是否包含单词word
func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])
	var dfs func(board [][]byte, i, j, k int) bool
	dfs = func(board [][]byte, i, j, k int) bool {
		if i >= row || i < 0 || j >= col || j < 0 || board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		board[i][j] = byte(' ')
		res := dfs(board, i, j+1, k+1) || dfs(board, i+1, j, k+1) || dfs(board, i, j-1, k+1) || dfs(board, i-1, j, k+1)
		board[i][j] = word[k]
		return res
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if dfs(board, i, j, 0) {
				return true
			}
		}
	}
	return false
}

// [0,0]-->[m,n]矩阵 坐标和不超过k，可以有走多少个格子
func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			visited[i][j] = false
		}
	}
	sum := func(i, j int) int {
		return i/10 + i%10 + j/10 + j%10
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= m || i < 0 || j >= n || j < 0 || visited[i][j] || sum(i, j) > k {
			return 0
		}
		visited[i][j] = true
		return dfs(i+1, j) + dfs(i-1, j) + dfs(i, j+1) + dfs(i, j-1) + 1
	}
	return dfs(0, 0)
}

// 查找二叉树和的路径
// dfs+回溯：先递归，target每次减去root的值，当左叶子节点不满足条件时，需要走右叶子结点，记住需要将路径中最后一个元素给去掉
func pathSum(root *TreeNode, target int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(root *TreeNode, target int)
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		target -= root.Val
		path = append(path, root.Val)
		defer func() { path = path[:len(path)-1] }()
		if root.Left == nil && root.Right == nil && target == 0 {
			ret = append(ret, append([]int(nil), path...))
			return
		}
		dfs(root.Left, target)
		dfs(root.Right, target)
	}
	dfs(root, target)
	return ret
}

// 二叉搜索树，第k大元素
func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	tree := make([]int, 0)
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		tree = append(tree, root.Val)
		traverse(root.Right)
	}
	traverse(root)
	return tree[len(tree)-k]
}
