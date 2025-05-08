package main

import "fmt"

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func (l *LinkedList[T]) InsertAtEnd(value T) {
	newNode := &Node[T]{data: value}

	if l.head == nil && l.tail == nil {
		l.head = newNode
		l.tail = newNode
		l.size = 1
		return
	}

	l.tail.next = newNode
	l.tail = newNode
	l.size++
}

func (l *LinkedList[T]) InsertAtBeginning(value T) {
	newNode := &Node[T]{data: value}

	if l.head == nil && l.tail == nil {
		l.head = newNode
		l.tail = newNode
		l.size = 1
		return
	}

	newNode.next = l.head
	l.head = newNode
	l.size++
}

func (l *LinkedList[T]) Delete(value T) bool {
	if l.head == nil || l.tail == nil {
		return false
	}

	if l.head.data == value {
		l.head = l.head.next
		l.size--
		return true
	}

	current := l.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
		l.size--
		return true
	}
	return false
}

func (l *LinkedList[T]) Display() {
	current := l.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
	fmt.Println("size ", l.size)
}

func main() {
	list := &LinkedList[int]{}
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtBeginning(30)
	list.InsertAtBeginning(40)
	list.InsertAtEnd(50)
	list.Delete(10)
	list.Display()
}
