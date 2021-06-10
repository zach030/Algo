package stack

import "log"

type ArrStack struct {
	stack []string
	len   int
	size  int
}

func NewArrStack(size int) *ArrStack {
	return &ArrStack{
		stack: make([]string, size),
		len:   0,
		size:  size,
	}
}

func (a *ArrStack) push(s string) bool {
	if a.len == a.size {
		log.Println("stack is full")
		return false
	}
	a.stack = append(a.stack, s)
	a.len = len(a.stack)
	return true
}

func (a *ArrStack) pop() string {
	if a.len == 0 {
		log.Println("empty stack")
		return ""
	}
	s := a.stack[a.len-1]
	a.stack = a.stack[:a.len-1]
	a.len--
	return s
}
