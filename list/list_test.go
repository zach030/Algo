package list

import (
	"fmt"
	"testing"
)

func TestList_Add(t *testing.T) {
	l := NewListWithHead()
	l.AddToLast(2)
	l.AddToLast(3)
	l.AddToLast(4)
	l.AddToLast(5)

	l.Print()

	l.MoveToFront(6)
	l.MoveToFront(2)
	l.Print()

	l.DeleteElement(3)

	l.Print()
}

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr!=nil{
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}

func TestReverseList(t *testing.T){
	h := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	res := reverseList(h)
	for res!=nil{
		fmt.Println(res.Val)
		res = res.Next
	}
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA==nil || headB==nil{
		return nil
	}
	startA,startB := headA,headB
	for headA!=headB{
		if headA!=nil{
			headA = headA.Next
		}else{
			headA = startB
		}
		if headB!=nil{
			headB = headB.Next
		}else{
			headB = startA
		}
	}
	return headA
}

func TestGetInsertNode(t *testing.T) {
	common := &ListNode{
		Val:  8,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	h1 := &ListNode{
		Val:  4,
		Next: &ListNode{
			Val:  1,
			Next: common,
		},
	}
	h2 := &ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  0,
			Next: &ListNode{
				Val:  1,
				Next: common,
			},
		},
	}
	ret := getIntersectionNode(h1,h2)
	for ret!=nil{
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}