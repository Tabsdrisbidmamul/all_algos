package main

import "testing"

func TestStackPeek(t *testing.T) {
	stack := Stack[int]{}

	stack.push(10)
	stack.push(20)
	stack.push(30)

	result, err := stack.peek()

	if err != nil {
		t.Fatalf("stack failed, %v", err)
	}

	if result != 30 {
		t.Errorf("expected %d, got %d", 30, result)
	}
}
