package daily

import (
	"log"
	"strings"
)

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

// 获取链表中倒数第k个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

// 合并两个有序的链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	n := &ListNode{Val: 0, Next: &ListNode{}}
	ret := n
	h1, h2 := l1, l2
	for h1 != nil && h2 != nil {
		var val int
		if h1.Val <= h2.Val {
			val = h1.Val
			h1 = h1.Next
		} else {
			val = h2.Val
			h2 = h2.Next
		}
		n.Next.Val = val
		n.Next.Next = &ListNode{}
		n = n.Next
	}
	if h1 == nil {
		n.Next = h2
	}
	if h2 == nil {
		n.Next = h1
	}
	return ret.Next
}

// 寻找两个链表的第一个公共节点
// 1.用map来存所有的节点信息，两个链表遍历
// 2.双指针遍历，两条路拼起来
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

// 使得所有奇数在数组的前半部分，所有偶数在数组的后半部分
// [1,2,3,4] ==> [1,3,2,4]
func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]%2 != 0 {
			left++
			continue
		}
		if nums[right]%2 == 0 {
			right--
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}

// 递增排序的数组，找和为target的组合
func twoSum(nums []int, target int) []int {
	var ret []int
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] > target {
			right--
			continue
		}
		if nums[left]+nums[right] < target {
			left++
			continue
		}
		return []int{nums[left], nums[right]}
	}
	return ret
}

// 反转单词顺序
// "the sky is blue" ===> "blue is sky the"
func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	s = strings.TrimLeft(s, " ")
	s = strings.TrimRight(s, " ")
	log.Println(strings.Split(s, " "))
	return ""
}
