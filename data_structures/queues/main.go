package main

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) enqueue(value T) {
	q.elements = append(q.elements, value)
}

func (q *Queue[T]) dequeue() (T, error) {
	var init_t T
	if q.is_empty() {
		return init_t, errors.New("queue is empty")
	}

	value := q.elements[0]
	q.elements = q.elements[1:]

	return value, nil
}

func (q *Queue[T]) size() int {
	return len(q.elements)
}

func (q *Queue[T]) is_empty() bool {
	return len(q.elements) == 0
}

func main() {
	q := &Queue[int]{}

	q.enqueue(10)
	q.enqueue(20)
	q.enqueue(30)

	fmt.Println("Queue size", q.size())
	item, err := q.dequeue()

	if err != nil {
		panic(err)
	}

	fmt.Println("dequeued item: ", item)
}
