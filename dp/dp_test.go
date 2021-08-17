package dp

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("babad"))
}

func TestFib(t *testing.T) {
	fmt.Println(fib(95))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestReverseList(t *testing.T) {
	l := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
					},
				},
			},
		},
	}
	k := 2
	i := 0
	for i < k {
		reverseN(l, k*i, (1+i)*k)
		i++
	}
	n := l
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

func reverseN(node *ListNode, l, r int) *ListNode {
	head := node
	left := node
	right := node
	i := 0
	for i < r {
		if i == l {
			left = head
		}
		head = head.Next
		i++
	}
	right = head
	var reverse func(node, tail *ListNode) *ListNode
	reverse = func(head, tail *ListNode) *ListNode {
		if head == nil {
			return nil
		}
		var prev *ListNode
		for head.Next != tail {
			tmp := head.Next
			head.Next = prev
			prev = head
			head = tmp
		}
		return prev
	}
	return reverse(left, right)
}
