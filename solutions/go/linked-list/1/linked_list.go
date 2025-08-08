package linkedlist

type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

func NewList(args ...interface{}) *List {
	if len(args) == 0 {
		return &List{}
	}
	list := List{}
	for _, v := range args {
		list.Push(v)
	}
	return &list
}

func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n == nil {
		return nil
	}
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	newHead := Node{Value: v}
	if l.head == nil {
		l.head = &newHead
		l.tail = &newHead
		return
	}
	oldHead := l.head
	oldHead.prev = &newHead
	newHead.next = oldHead
	l.head = &newHead
}

func (l *List) Push(v interface{}) {
	newTail := Node{Value: v}
	if l.tail == nil {
		l.tail = &newTail
		l.head = &newTail
		return
	}
	oldTail := l.tail
	oldTail.next = &newTail
	newTail.prev = oldTail
	l.tail = &newTail
}

func (l *List) Shift() (interface{}, error) {
	value := l.head.Value
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return value, nil
	}
	l.head = l.head.next
	l.head.prev = nil
	return value, nil
}

func (l *List) Pop() (interface{}, error) {
	value := l.tail.Value
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return value, nil
	}
	l.tail = l.tail.prev
	l.tail.next = nil
	return value, nil
}

func (l *List) Reverse() {
	n := l.head
	for n != nil {
		next := n.next
		n.next, n.prev = n.prev, n.next
		n = next
	}
	l.head, l.tail = l.tail, l.head
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
