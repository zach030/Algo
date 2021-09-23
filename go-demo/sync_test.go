package go_demo

import (
	"fmt"
	"testing"
)

func Test_syncDemo(t *testing.T) {
	syncDemo()
}

type ListNode struct {
	val  int
	next *ListNode
}

func createList(n int) *ListNode {
	head := &ListNode{val: 0, next: &ListNode{}}
	h := head
	head = head.next
	for i := 1; i <= n; i++ {
		head.val = i
		head.next = &ListNode{}
		if i == n {
			head.next = nil
		}
		head = head.next
	}
	return h.next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v ", head.val)
		head = head.next
	}
	fmt.Println()
}

func reverseN(head *ListNode, n int) *ListNode {
	var last *ListNode
	var reverse func(head *ListNode, n int) *ListNode
	reverse = func(head *ListNode, n int) *ListNode {
		if head == nil || head.next == nil {
			return head
		}
		if n == 0 {
			return head
		}
		if n == 1 {
			last = head.next
			return head
		}
		lastHead := reverse(head.next, n-1)
		head.next.next = head
		head.next = last
		return lastHead
	}
	return reverse(head, n)
}

func TestReverse(t *testing.T) {
	for i := 0; i <= 7; i++ {
		fmt.Println(i)
		head := createList(6)
		printList(head)
		h := reverseN(head, i)
		printList(h)
	}

	//head := createList(6)
	//printList(head)
	//h := reverseList(head, 7)
	//printList(h)
}
