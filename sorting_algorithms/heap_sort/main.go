package main

import "fmt"

func heapify(items []int, n, index int) {
	// initialise largest as root
	largest := index
	left := (index << 1) + 1
	right := (index << 1) + 2

	if left < n && items[left] > items[largest] {
		largest = left
	}

	if right < n && items[right] > items[largest] {
		largest = right
	}

	fmt.Printf("heapify: largest %d\n", largest)
	fmt.Printf("heapify: index %d\n", index)

	if largest != index {
		items[index], items[largest] = items[largest], items[index]
		fmt.Printf("heapify: array %v\n", items)
		heapify(items, n, largest)
	}
}

func heap_sort(items []int) {
	n := len(items)

	for i := n >> 1; i >= 0; i-- {
		fmt.Printf("heap_sort 1st: i %d\n", i)
		heapify(items, n, i)
	}

	for i := n - 1; i > 0; i-- {
		fmt.Printf("heap_sort 2nd: i %d\n", i)
		// swap max element for root
		items[0], items[i] = items[i], items[0]
		heapify(items, i, 0)
	}
}

func main() {
	items := []int{4, 10, 3, 5, 1}
	fmt.Println("Unsorted array:", items)

	heap_sort(items)

	fmt.Println("Sorted array:", items)
}
