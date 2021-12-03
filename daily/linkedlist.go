package daily

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转打印链表
func reversePrint(head *ListNode) []int {
	tmp := make([]int, 0)
	var recur = func(node *ListNode) {}
	recur = func(node *ListNode) {
		if node == nil {
			return
		}
		recur(node.Next)
		tmp = append(tmp, node.Val)
	}
	recur(head)
	return tmp
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 拷贝复杂链表
func copyRandomList(head *Node) *Node {
	store := make(map[*Node]*Node)
	curr := head
	for curr != nil {
		store[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}
	curr = head
	for curr != nil {
		store[curr].Next = store[curr.Next]
		store[curr].Random = store[curr.Random]
		curr = curr.Next
	}
	return store[head]
}
