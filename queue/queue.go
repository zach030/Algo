package queue

import (
	"Algo/list"
	"log"
)

type Queue interface {
	EnQueue(string2 string) bool
	DeQueue() string
}

// 顺序队列
type SeqQueue struct {
	items []string
	size  int
	head  int
	tail  int
}

func NewSeqQueue(size int) *SeqQueue {
	return &SeqQueue{
		items: make([]string, size),
		size:  size,
		head:  0,
		tail:  0,
	}
}

func (s *SeqQueue) EnQueue(str string) bool {
	if s.tail == s.size {
		if s.head == 0 {
			log.Println("queue is full")
			return false
		}
		for i := s.head; i < s.tail; i++ {
			s.items[i-s.head] = s.items[i]
		}
		s.tail -= s.head
		s.head = 0
	}
	s.items[s.tail] = str
	s.tail++
	return true
}

func (s *SeqQueue) DeQueue() string {
	if s.head == s.tail {
		log.Println("queue is empty")
		return ""
	}
	ele := s.items[s.head]
	s.head++
	return ele
}

// 链式队列
type LinkedQueue struct {
	head *list.Node
	tail *list.Node
}

func NewLinkedQueue()*LinkedQueue{
	lq := &LinkedQueue{}
	return lq
}

// 循环队列
type RingQueue struct {

}