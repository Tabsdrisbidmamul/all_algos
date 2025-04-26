package main

import (
	"reflect"
	"testing"
)

func TestBucketSort(t *testing.T) {
	tests := []struct {
		input, expected []float64
	}{
		{
			[]float64{0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68},
			[]float64{0.12, 0.17, 0.21, 0.23, 0.26, 0.39, 0.68, 0.72, 0.78, 0.94},
		},
		{
			input:    []float64{0.5, 0.4, 0.3, 0.2, 0.1},
			expected: []float64{0.1, 0.2, 0.3, 0.4, 0.5},
		},
		{
			input:    []float64{},
			expected: []float64{},
		},
	}

	for _, test := range tests {
		result := bucket_sort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Expected %v\nGot %v", test.expected, result)
		}
	}
}
