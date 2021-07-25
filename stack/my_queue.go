package stack

// 使用两个栈实现先入先出队列
type MyQueue struct {
	in  []int
	out []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

/** Push element x to the back of in. */
func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

/** Removes the element from in front of in and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.out) == 0 {
		// move from in to out
		for i := len(this.in) - 1; i >= 0; i-- {
			this.out = append(this.out, this.in[i])
			this.in = this.in[:len(this.in)-1]
		}
	}
	out := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return out
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.out) == 0 {
		// move from in to out
		for i := len(this.in) - 1; i >= 0; i-- {
			this.out = append(this.out, this.in[i])
			this.in = this.in[:len(this.in)-1]
		}
	}
	return this.out[len(this.out)-1]
}

/** Returns whether the in is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.in)==0 && len(this.out)==0
}
