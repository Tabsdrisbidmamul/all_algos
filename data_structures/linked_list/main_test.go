package main

import "testing"

func TestInsertAtBeginning(t *testing.T) {
	list := &LinkedList[int]{}
	list.InsertAtBeginning(20)
	list.InsertAtBeginning(10)

	if list.head.data != 10 && list.size != 2 {
		t.Errorf("expected first element %d and size %d.\nGot %d and %d", 10, 2, list.head.data, list.size)
	}
}

func TestInsertAtEnd(t *testing.T) {
	list := &LinkedList[int]{}
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)

	if list.tail.data != 20 && list.size != 2 {
		t.Errorf("expected first element %d and size %d.\nGot %d and %d", 20, 1, list.tail.data, list.size)
	}
}

func TestDelete(t *testing.T) {
	list := &LinkedList[int]{}
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtEnd(40)

	if list.tail.data != 40 && list.size != 4 {
		t.Errorf("expected first element %d and size %d.\nGot %d and %d", 40, 4, list.tail.data, list.size)
	}

	list.Delete(20)

	if list.size != 3 {
		t.Errorf("expected size %d.\nGot %d", 3, list.size)
	}

	current := list.head
	if current.data != 10 {
		t.Errorf("expected element %d.\nGot %d", 10, current.data)
	}

	current = current.next
	if current.data != 30 {
		t.Errorf("expected element %d.\nGot %d", 30, current.data)
	}

	current = current.next
	if current.data != 40 {
		t.Errorf("expected element %d.\nGot %d", 40, current.data)
	}
}
