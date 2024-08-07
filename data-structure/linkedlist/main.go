package main

import "fmt"

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) First() *Node {
	return l.head
}

func (l *LinkedList) Last() *Node {
	return l.tail
}

func (l *LinkedList) Push(value int) {
	var node *Node = &Node{
		value: value,
	}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node

}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) Next() *Node {
	return n.next
}

func main() {
	l := &LinkedList{}
	l.Push(1)
	l.Push(2)
	l.Push(3)

	// fmt.Printf("first: %v\n", l.First())
	// fmt.Printf("next: %v\n", l.head.next)
	// fmt.Printf("tail: %v\n", l.tail)

	n := l.First()
	for {
		fmt.Printf("%v\n", n.value)
		n = n.Next()
		if n == nil {
			break
		}
	}

	p := l.Last()
	for {
		fmt.Printf("%v\n", p.value)
		p = p.Prev()
		if p == nil {
			break
		}
	}
}
