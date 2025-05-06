package main

import (
	"testing"
)

func TestEnqueueDequeue(t *testing.T) {
	queue := NewQueue[int](2)

	if err := queue.Enqueue(1); err != nil {
		t.Fatalf("unexpected error on enqueue: %v", err)
	}

	if err := queue.Enqueue(2); err != nil {
		t.Fatalf("unexpected error on enqueue: %v", err)
	}

	val1, err := queue.Dequeue()
	if err != nil || val1 != 1 {
		t.Errorf("expected %d. got %d with err %v", 1, val1, err)
	}

	val2, err := queue.Dequeue()
	if err != nil || val2 != 2 {
		t.Errorf("expected %d. got %d with err %v", 2, val2, err)
	}
}

func TestQueueBlocking(t *testing.T) {
	queue := NewQueue[int](1)
	done := make(chan bool)

	go func() {
		queue.Enqueue(1)
		queue.Enqueue(2)
		done <- true
	}()

	val1, _ := queue.Dequeue()
	if val1 != 1 {
		t.Errorf("expected %d, got %d", 1, val1)
	}

	val2, _ := queue.Dequeue()
	if val2 != 2 {
		t.Errorf("expected %d, got %d", 2, val2)
	}

	<-done
}

func TestQueueClose(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic when sending to a closed channel, but did not panic")
		}
	}()

	queue := NewQueue[int](1)
	queue.Close()

	_, _ = queue.Dequeue()
	queue.Enqueue(1)
}
