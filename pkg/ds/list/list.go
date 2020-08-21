package list

// 线程安全

type Node struct {
	data interface{}
	next *Node
}

type List struct {
	head *Node
	tail *Node
	len  int
}

func NewList() *List {
	l := new(List)
	l.head = new(Node)
	return l
}

func (list *List) Push(e interface{}) {
	node := new(Node)
	node.data = e
	node.next = list.head.next

	list.head.next = node
	list.len++
}

func (list *List) Length() int {
	return list.len
}
