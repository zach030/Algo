package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// 回文链表判断
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	iterator := head
	arr := make([]int, 0)
	for iterator != nil {
		arr = append(arr, iterator.Val)
		iterator = iterator.Next
	}
	left, right := 0, len(arr)-1
	for left < right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 删除链表的倒数第 n 个结点，并且返回链表的头结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

// 判断环形链表
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			h := head
			for h != slow {
				h = h.Next
				slow = slow.Next
			}
			return h
		}
	}
	return nil
}

// 翻转部分链表
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	leftNode := pre.Next
	currNode := rightNode.Next

	pre.Next = nil
	rightNode.Next = nil

	reverseList(leftNode)

	pre.Next = rightNode
	leftNode.Next = currNode
	return dummyNode.Next
}

// 翻转全部链表
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
