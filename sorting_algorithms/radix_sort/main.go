package main

import "fmt"

func get_max(arr []int) int {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	return max
}

// Stable sort on exp (10^i) where i is the digit pos, 1's, 10's, 100's
func counting_sort(arr []int, exp int) []int {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	for i := 0; i < n; i++ {
		digit := (arr[i] / exp) % 10
		count[digit]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / exp) % 10
		output[count[digit]-1] = arr[i]
		count[digit]--
	}

	return output
}

func radix_sort(arr []int) []int {
	max := get_max(arr)

	for exp := 1; max/exp > 0; exp *= 10 {
		arr = counting_sort(arr, exp)
	}

	return arr
}

func main() {
	arr := []int{170, 45, 75, 90, 802, 24, 2, 66}
	fmt.Println("Unsorted array:", arr)
	sortedArr := radix_sort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
