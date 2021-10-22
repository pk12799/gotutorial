package main

import (
	"fmt"
)

type Node struct {
	next *Node
	prev *Node
	key  interface{}
}
type List struct {
	head *Node
	tail *Node
}

func (L *List) insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list
	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}
func (l *List) display() {
	list := l.head
	for list != nil {
		fmt.Printf("--> %+v ", list.key)
		list = list.next
	}
	fmt.Println()
}
func display(List *Node) {

	for List != nil {
		fmt.Printf("-->%v", List.key)
		List = List.next
	}
	fmt.Println()
}

func showBack(list *Node) {
	for list != nil {
		fmt.Printf("<--%v", list.key)
		list = list.prev
	}
	fmt.Println()
}

func (l *List) reverse() {
	cur := l.head
	var prev *Node
	l.tail = l.head
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	l.head = prev
	display(l.head)
}

// type e interface{}

func main() {
	link := &List{}
	n := 0
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		// f := bufio.NewScanner(os.Stdin)
		// for f.Scan() {
		// 	input := f.Text()
		// 	link.insert(input)
		// 	link.display()
		// }
		var e int
		fmt.Scan(&e)
		link.insert(e)

	}
	fmt.Println("----------------------------------")
	fmt.Printf("head : %v\n", link.head.key)
	fmt.Printf("Tail : %v\n", link.tail.key)
	link.display()
	fmt.Println("\n==============================\n")
	fmt.Printf("head: %v\n", link.head.key)
	fmt.Printf("tail: %v\n", link.tail.key)
	link.reverse()
	fmt.Println("\n==============================\n")

}
