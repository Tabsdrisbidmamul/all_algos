package main

import "testing"

func TestEnqueue(t *testing.T) {
	q := &Queue[int]{}

	if q.size() != 0 {
		t.Fatalf("expected %d, got %d", 0, q.size())
	}

	q.enqueue(10)

	if q.size() != 1 {
		t.Errorf("expected %d, got %d", 1, q.size())
	}
}

func TestDequeue(t *testing.T) {
	q := &Queue[int]{}

	if q.size() != 0 {
		t.Fatalf("expected %d, got %d", 0, q.size())
	}

	q.enqueue(10)
	q.enqueue(20)
	q.enqueue(30)

	if q.size() != 3 {
		t.Errorf("expected %d, got %d", 3, q.size())
	}

	item, err := q.dequeue()
	if err != nil {
		t.Fatalf("expected %d, got %v", 30, err)
	}

	if q.size() != 2 && item != 30 {
		t.Errorf("expected queue size %d and item %d. Got %d and %d", 2, 30, q.size(), item)
	}
}
