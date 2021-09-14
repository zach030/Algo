package queue

import (
	"container/list"
)

// MonotonicQueue 单调队列
type MonotonicQueue struct {
	l *list.List
}

func (m *MonotonicQueue) push(n int) {
	for m.l.Len() != 0 && m.l.Back().Value.(int) < n {
		m.l.Remove(m.l.Back())
	}
	m.l.PushBack(n)
}

func (m *MonotonicQueue) max() int {
	return m.l.Front().Value.(int)
}

func (m *MonotonicQueue) pop(n int) {
	if n == m.l.Front().Value.(int) {
		m.l.Remove(m.l.Front())
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	window := &MonotonicQueue{l: list.New()}
	ret := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			window.push(nums[i])
			continue
		}
		window.push(nums[i])
		ret = append(ret, window.max())
		window.pop(nums[i-k+1])
	}
	return ret
}
