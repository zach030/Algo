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

func Test_exchange(t *testing.T) {
	t1 := []int{1, 2, 3, 4}
	t2 := []int{2, 4, 3, 1}
	t3 := []int{1, 3, 4, 2}

	fmt.Println(exchange(t1))
	fmt.Println(exchange(t2))
	fmt.Println(exchange(t3))
}

func Test_reverseWords(t *testing.T) {
	words := "  the   sky is blue"
	fmt.Println(reverseWords(words))
}
