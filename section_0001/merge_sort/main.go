package main

import "fmt"

func merge_sort(items []int) []int {
	if len(items) <= 1 {
		return items
	}

	// left is from 0 and up to and not including mid point
	// right is from mid till the end of array
	left := merge_sort(items[:len(items)>>1])
	right := merge_sort(items[len(items)>>1:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}

	fmt.Printf("unsorted %v\n", unsorted)
	sorted := merge_sort(unsorted)
	fmt.Printf("sorted %v\n", sorted)
}
