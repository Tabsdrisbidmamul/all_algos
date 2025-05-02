package main

import "testing"

func TestStackPeek(t *testing.T) {
	stack := Stack[int]{}

	stack.push(10)
	stack.push(20)
	stack.push(30)

	result, err := stack.peek()

	if err != nil {
		t.Fatalf("peek failed, expected panic, got %v", err)
	}

	if result != 30 {
		t.Errorf("expected %d, got %d", 30, result)
	}
}

func TestStackPush(t *testing.T) {
	stack := Stack[int]{}

	_, err := stack.peek()

	if err == nil {
		t.Fatalf("peek failed, expected panic, got %v", err)
	}

	stack.push(10)

	result, err := stack.peek()

	if err != nil {
		t.Fatalf("peek failed, expected %d, got %v", 10, err)
	}

	if result != 10 {
		t.Errorf("expected %d, got %d", 10, result)
	}

	if len(stack.elements) != 1 {
		t.Errorf("expected %d, got %d", 1, len(stack.elements))
	}
}

func TestStackPop(t *testing.T) {
	stack := Stack[int]{}

	stack.push(10)
	stack.push(20)
	stack.push(30)

	result, err := stack.pop()

	if err != nil {
		t.Fatalf("peek failed, expected panic, got %v", err)
	}

	if result != 30 {
		t.Errorf("expected %d, got %d", 30, result)
	}

	if len(stack.elements) != 2 {
		t.Errorf("expected %d, got %d", 2, len(stack.elements))
	}
}
