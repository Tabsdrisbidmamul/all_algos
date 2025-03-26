package main

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		unsorted_array, expected []int
	}{
		{
			unsorted_array: []int{5, 2, 4, 6, 1, 3},
			expected:       []int{1, 2, 3, 4, 5, 6},
		},
		{
			unsorted_array: []int{31, 41, 59, 26, 41, 58},
			expected:       []int{26, 31, 41, 41, 58, 59},
		},
		{
			unsorted_array: []int{2, 1},
			expected:       []int{1, 2},
		},
		{
			unsorted_array: []int{1, 2, 3, 4, 5},
			expected:       []int{1, 2, 3, 4, 5},
		},
		{
			unsorted_array: []int{},
			expected:       []int{},
		},
	}

	for _, test := range tests {
		result := insertion_sort(test.unsorted_array, false)
		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("Expected: %v\nGot: %v", test.expected, result)
		}
	}
}

func TestInsertionSortReverse(t *testing.T) {
	tests := []struct {
		unsorted_array, expected []int
	}{
		{
			unsorted_array: []int{5, 2, 4, 6, 1, 3},
			expected:       []int{6, 5, 4, 3, 2, 1},
		},
		{
			unsorted_array: []int{31, 41, 59, 26, 41, 58},
			expected:       []int{59, 58, 41, 41, 31, 26},
		},
		{
			unsorted_array: []int{2, 1},
			expected:       []int{2, 1},
		},
		{
			unsorted_array: []int{},
			expected:       []int{},
		},
	}

	for _, test := range tests {
		result := insertion_sort_reverse(test.unsorted_array)
		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("Expected: %v\nGot: %v", test.expected, result)
		}
	}
}
