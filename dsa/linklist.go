package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type Linkl struct {
	head   *Node
	length int
}

func (l *Linkl) Insert(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}
func (l Linkl) display() {
	top := l.head
	for l.length != 0 {
		fmt.Println(top.data)
		top = top.next
		l.length--
	}

}

func (l *Linkl) delete(val int) {
	if l.length == 0 {
		return
	}
	if l.head.data == val {
		l.head = l.head.next
		l.length--

	}
	prev := l.head
	for prev.next.data != val {
		prev = prev.next.next
	}
	prev.next = prev.next.next
	l.length--
}

func main() {
	list := &Linkl{}
	node1 := &Node{data: 45}
	node3 := &Node{data: 78}
	node2 := &Node{data: 21}

	node4 := &Node{data: 58}

	list.Insert(node1)

	node5 := &Node{data: 48}

	list.Insert(node2)

	node6 := &Node{data: 11}

	list.Insert(node3)
	list.Insert(node6)
	list.Insert(node5)
	list.Insert(node4)
	fmt.Println(*list.head)
	list.display()
	list.delete(48)
	list.delete(11)
	// list.delete(21)
	// list.delete(78)

	fmt.Println(*list.head)
	// list.delete(58)
	fmt.Println(*list.head)
	list.display()
}
