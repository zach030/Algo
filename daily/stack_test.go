package daily

import (
	"fmt"
	"testing"
)

//["MinStack","push","push","push","top","pop","min","pop","min","pop","push","top","min","push","top","min","pop","min"]
//[[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]
func TestMinStack_Push(t *testing.T) {
	m := Constructor()
	m.Push(2147483646)
	m.Push(2147483646)
	m.Push(2147483647)
	fmt.Println("top:", m.Top())
	m.Pop()
	fmt.Println("min:", m.Min())
	m.Pop()
	fmt.Println("min:", m.Min())
	m.Pop()
	m.Push(2147483647)
	fmt.Println("top:", m.Top())
	fmt.Println("min:", m.Min())
	m.Push(-2147483648)
	fmt.Println("top:", m.Top())
	fmt.Println("min:", m.Min())
	m.Pop()
	fmt.Println("min:", m.Min())
}

//[null,null,null,null,2147483647,null,2147483646,null,2147483646,null,null,2147483647,2147483647,null,-2147483648,-2147483648,null,2147483647]
