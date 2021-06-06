package list

import (
	"testing"
)

func TestList_Add(t *testing.T) {
	l := NewListWithHead(Val{Value: 1})
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
