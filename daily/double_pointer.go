package daily

// 删除链表中的指定节点
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	prev, curr := head, head.Next
	for curr != nil && curr.Val != val {
		prev = curr
		curr = curr.Next
	}
	if curr != nil {
		prev.Next = curr.Next
	}
	return head
}
