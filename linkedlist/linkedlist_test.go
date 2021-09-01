package linkedlist

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(isPalindrome(list))
}

func TestHasCycle(t *testing.T) {
	common := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	common.Next.Next.Next = common
	h := &ListNode{
		Val:  3,
		Next: common,
	}
	fmt.Println(detectCycle(h))
}

func TestReverseList(t *testing.T) {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	res := reverseList(h)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func TestReverseBetween(t *testing.T) {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	res := reverseBetween(h, 2, 4)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func TestMergeList(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{2, &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{2, &ListNode{Val: 3}}}
	l := mergeTwoLists(l1, l2)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
