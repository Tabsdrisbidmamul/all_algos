package main

import (
	"sync"
	"testing"
)

func TestConcurrentPush(t *testing.T) {
	stack := &Stack[int]{}
	var wg sync.WaitGroup

	const num_goroutines = 10
	const num_ops_per_goroutine = 100

	for i := 0; i < num_goroutines; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			for j := 0; j < num_ops_per_goroutine; j++ {
				stack.push(start*num_ops_per_goroutine + j)
			}
		}(i)
	}

	wg.Wait()

	if stack.length() != num_goroutines*num_ops_per_goroutine {
		t.Errorf("expected stack length %d, got %d", num_goroutines*num_ops_per_goroutine, stack.length())
	}
}

func TestConcurrentPop(t *testing.T) {
	stack := &Stack[int]{}
	var wg sync.WaitGroup
	var guard sync.Mutex

	const num_goroutines = 10
	const num_ops_per_goroutine = 10
	results := make(map[int]bool, num_goroutines*num_ops_per_goroutine)

	for i := 0; i < num_goroutines*num_ops_per_goroutine; i++ {
		stack.push(i)
	}

	for i := 0; i < num_goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < num_ops_per_goroutine; j++ {
				value, err := stack.pop()
				if err != nil {
					t.Errorf("stack is empty at %d %d", i, j)
				}

				guard.Lock()
				results[value] = true
				guard.Unlock()
			}
		}()
	}

	wg.Wait()

	if len(results) != num_goroutines*num_ops_per_goroutine {
		t.Errorf("expected to pop %d, got %d", num_goroutines*num_ops_per_goroutine, len(results))
	}
}
