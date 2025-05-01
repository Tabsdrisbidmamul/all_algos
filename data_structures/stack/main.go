package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(value T) {
	s.elements = append(s.elements, value)
}

func (s *Stack[T]) pop() (T, error) {
	var init_t T
	if s.is_empty() {
		return init_t, errors.New("stack is empty")
	}

	last_index := len(s.elements) - 1
	value := s.elements[last_index]
	s.elements = s.elements[:last_index]
	return value, nil
}

func (s *Stack[T]) peek() (T, error) {
	var init_t T
	if s.is_empty() {
		return init_t, errors.New("stack is empty")
	}

	return s.elements[len(s.elements)-1], nil

}

func (s *Stack[T]) is_empty() bool {
	return len(s.elements) == 0
}

func main() {
	stack := Stack[int]{}

	stack.push(10)
	stack.push(20)
	stack.push(30)

	fmt.Println("Top element: ", must(stack.peek()))
}

func must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}

	return val
}
