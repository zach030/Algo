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
	if head==nil || head.Next==nil{
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
