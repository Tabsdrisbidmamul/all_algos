package main

import (
	"errors"
	"fmt"
	"sync"
)

type Stack[T any] struct {
	guard    sync.Mutex
	elements []T
}

func (s *Stack[T]) push(value T) {
	s.guard.Lock()
	defer s.guard.Unlock()

	s.elements = append(s.elements, value)
}

func (s *Stack[T]) peek() (T, error) {
	s.guard.Lock()
	defer s.guard.Unlock()

	var init_t T
	if len(s.elements) == 0 {
		return init_t, errors.New("stack is empty")
	}

	return s.elements[len(s.elements)-1], nil
}

func (s *Stack[T]) pop() (T, error) {
	s.guard.Lock()
	defer s.guard.Unlock()

	var init_t T
	if len(s.elements) == 0 {
		return init_t, errors.New("stack is empty")
	}

	last_index := len(s.elements) - 1
	last_element := s.elements[last_index]
	s.elements = s.elements[:last_index]
	return last_element, nil
}

func (s *Stack[T]) length() int {
	s.guard.Lock()
	defer s.guard.Unlock()
	return len(s.elements)
}

func main() {
	stack := &Stack[int]{}
	stack.push(1)
	stack.push(2)

	top, err := stack.pop()

	if err != nil {
		panic("stack is empty")
	}

	fmt.Println("Popped:", top)
}
