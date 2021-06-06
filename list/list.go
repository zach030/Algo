package list

import (
	"fmt"
)

type Method interface {
	Insert(i int,val interface{})
	AddToLast(val interface{})
	DeleteWithPos(i int)
	DeleteElement(i interface{})
	Len()int

	Print()
	MoveToFront(i interface{})
	Back()interface{}
}

type List struct {
	Head *Node
	Length int
}

func NewListWithHead(ele element)*List {
	return &List{
		Head: NewNode(ele),
		Length: 1,
	}
}

func NewList()*List {
	return &List{
		Head: &Node{Ele: Val{}},
		Length: 0,
	}
}

// 插入到原来第i个节点之前,成为第i个节点
func (l *List) Insert(i int, val interface{}) {
	n := NewNode(Val{Value: val})
	head := l.Head
	// 改头部元素
	if i==0{
		n.Next = head
		l.Length++
		l.Head = n
		return
	}
	for count := 0;count<=i;count++{
		// 找到前驱节点
		if count == i-1{
			n.Next = head.Next
			head.Next = n
			l.Length++
		}
		head = head.Next
	}
}

func (l *List) AddToLast(val interface{}) {
	l.Insert(l.Length,val)
}

// 删除第i个节点
func (l *List) DeleteWithPos(i int) {
	head := l.Head
	for count :=0;count<=i-1;count++{
		prev := head.Next
		if count==i-1{
			head.Next = prev.Next
			l.Length--
		}
		head = head.Next
	}
}

func (l *List) DeleteElement(i interface{}) {
	head := l.Head
	for head!=nil{
		prev := head.Next
		if prev.Ele.Get()==i{
			head.Next = prev.Next
		}
		head = head.Next
	}
}

func (l *List) Len() int {
	return l.Length
}

func (l *List) Print() {
	head := l.Head
	for head!=nil{
		fmt.Printf("%+v",head.Ele.Get())
		head = head.Next
	}
	fmt.Println()
}

// 将节点移到队列头
func (l *List) MoveToFront(val interface{}) {
	// 先找到此数据对应的节点
	head := l.Head
	pos := 0
	ele := Val{Value: val}
	for head.Ele!=ele{
		if pos==l.Length-1{
			// 此节点不存在，add
			l.Insert(0,val)
			return
		}
		pos++
		head = head.Next
	}
	// 删除此节点
	l.DeleteWithPos(pos)
	// 再移到队头
	l.Insert(0,val)
}

func (l *List) Back() interface{} {
	head := l.Head
	for head.Next!=nil{
		head = head.Next
	}
	return head.Ele.Get()
}