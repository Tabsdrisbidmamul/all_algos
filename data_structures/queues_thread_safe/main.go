package main

import (
	"fmt"
	"time"
)

type Queue[T any] struct {
	dataChan chan T
	done     chan struct{}
}

func NewQueue[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		dataChan: make(chan T, capacity),
		done:     make(chan struct{}),
	}
}

// only blocks when full or empty
func (q *Queue[T]) Enqueue(item T) error {
	select {
	case q.dataChan <- item:
		return nil
	case <-q.done:
		return fmt.Errorf("queue closed")
	}
}

func (q *Queue[T]) Dequeue() (T, error) {
	var init_t T
	select {
	case item := <-q.dataChan:
		return item, nil
	case <-q.done:
		return init_t, fmt.Errorf("queue closed")
	}
}

func (q *Queue[T]) Close() {
	close(q.done)
	close(q.dataChan)
}

func main() {
	queue := NewQueue[int](5)

	go func() {
		for i := 0; i < 10; i++ {
			if err := queue.Enqueue(i); err != nil {
				fmt.Println("Enqueue error", err)
				return
			}

			fmt.Println("Enqueued: ", i)
			time.Sleep(100 * time.Millisecond)
		}
		queue.Close()
	}()

	for {
		item, err := queue.Dequeue()
		if err != nil {
			fmt.Println("Dequeue error: ", err)
			break
		}

		fmt.Println("Dequeued: ", item)
		time.Sleep(1000 * time.Millisecond)
	}
}
