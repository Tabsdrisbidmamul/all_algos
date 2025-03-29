package main

import "fmt"

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quick_sort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quick_sort(arr, low, p-1)
		arr = quick_sort(arr, p+1, high)
	}
	return arr
}

func main() {
	arr := []int{1, 5, 2, 6}

	fmt.Printf("Unsorted: %v\n", arr)
	arr = quick_sort(arr, 0, len(arr)-1)
	fmt.Printf("Sorted: %v\n", arr)
}
