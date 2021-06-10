package stack

import (
	"fmt"
	"testing"
)

func TestArrStack_push(t *testing.T) {
	s := NewArrStack(5)
	s.push("a")
	s.push("b")
	s.push("c")
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
}
