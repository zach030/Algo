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

// 递归法
func reverseList2(head *ListNode) *ListNode {
	if head.Next == nil || head == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

var successor *ListNode

// 翻转前N个节点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween2(head.Next, left-1, right-1)
	return head
}

// k个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

// 翻转a-b之间的链表
func reverse(a, b *ListNode) *ListNode {
	var prev *ListNode
	curr, nxt := a, a
	for curr != b {
		nxt = curr.Next
		curr.Next = prev
		prev = curr
		curr = nxt
	}
	return prev
}

// 合并有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{Val: 0}
	head := l
	for l1 != nil && l2 != nil {
		head.Next = &ListNode{}
		head = head.Next
		if l1.Val <= l2.Val {
			head.Val = l1.Val
			l1 = l1.Next
		} else {
			head.Val = l2.Val
			l2 = l2.Next
		}
	}
	if l1 == nil {
		head.Next = l2
	}
	if l2 == nil {
		head.Next = l1
	}
	return l.Next
}

// 找出公共节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
			continue
		}
		if b == nil {
			b = headA
			continue
		}
		a = a.Next
		b = b.Next
	}
	return a
}
