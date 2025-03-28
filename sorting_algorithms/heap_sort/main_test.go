package main

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	tests := []struct {
		input, expected []int
	}{
		{
			input:    []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			input:    []int{4, 10, 3, 5, 1},
			expected: []int{1, 3, 4, 5, 10},
		},
		{
			input:    []int{},
			expected: []int{},
		},
	}

	for _, test := range tests {
		heap_sort(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("Expected %v\nGot %v", test.expected, test.input)
		}
	}
}
