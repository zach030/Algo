package daily

type MinStack struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	c := MinStack{
		data: make([]int, 0),
		min:  make([]int, 0),
	}
	return c
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.min) == 0 {
		this.min = append(this.min, x)
		return
	}
	peek := this.min[len(this.min)-1]
	if x <= peek {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.data) >= 1 {
		peek := this.data[len(this.data)-1]
		this.data = this.data[:len(this.data)-1]
		if peek == this.min[len(this.min)-1] {
			this.min = this.min[:len(this.min)-1]
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}
	return -1
}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
