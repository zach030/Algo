package list

type element interface {
	Set(interface{})
	Get()interface{}
}

type Val struct {
	Value interface{}
}

func (v Val) Set(i interface{}) {
	v.Value = i
}

func (v Val) Get() interface{} {
	return v.Value
}

// 有序单链表
type Node struct {
	Ele element
	Next *Node
}

func NewNode(ele element)*Node{
	return &Node{
		Ele:  ele,
		Next: nil,
	}
}