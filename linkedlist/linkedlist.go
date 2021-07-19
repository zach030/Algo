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
