package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		unsorted_array, expected []int
	}{
		{
			unsorted_array: []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9},
			expected:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			unsorted_array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			unsorted_array: []int{5, 4, 3, 2, 1},
			expected:       []int{1, 2, 3, 4, 5},
		},
		{
			unsorted_array: []int{},
			expected:       []int{},
		},
	}

	for _, test := range tests {
		result := merge_sort(test.unsorted_array)
		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("Expected: %v\nGot: %v", test.expected, result)
		}
	}
}
