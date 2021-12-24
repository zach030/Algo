package daily

import (
	"fmt"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	l := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5}}}
	h := deleteNode(l, 1)
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

func TestLastKNode(t *testing.T) {
	l := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}
	h := getKthFromEnd(l, 2)
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

func TestMergeList(t *testing.T) {
	l1 := &ListNode{Val: 3, Next: &ListNode{Val: 6, Next: &ListNode{Val: 9}}}
	l2 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: &ListNode{Val: 8}}}}
	h := mergeTwoLists(l1, l2)
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}
