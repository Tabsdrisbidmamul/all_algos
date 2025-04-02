package main

import "fmt"

func counting_sort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// find the maximum value in the array
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	// initialise the count array
	count := make([]int, max+1)

	for _, num := range arr {
		count[num]++
	}

	sorted_index := 0
	for i, c := range count {
		for c > 0 {
			arr[sorted_index] = i
			sorted_index++
			c--
		}
	}

	return arr
}

func main() {
	arr := []int{4, 2, 2, 8, 3, 3, 1}

	fmt.Printf("Original array: %v\n", arr)
	sortedArr := counting_sort(arr)
	fmt.Printf("Sorted array: %v\n", sortedArr)
}
