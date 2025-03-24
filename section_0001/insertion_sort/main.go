package main

import "fmt"

// run `go run main.go` to see debug run
// Insertion sort works like this:
// arr[0] is our sorted pile, and arr[1:n] is unsorted
// We have 2 pointers (double loop), the i pointer is the next element we look at, and we check to see if it can slot into our sorted pile.
// First iteration we have 5 as our sorted and 2 as new element to put into the sorted. 2 < 5, so we place 2 as the first element and 5 as the next element. We do this for each item.
// BEFORE [5 2 4 6 1 3]

// (i := 1 | values[i] := 2) key := 2  (j := 0 | values[j] := 5)
// doing a swap: values[j+1](2) := values[j](5)
// swapped iteration [5 5 4 6 1 3]
// values[j+1](5) = key(2)
// sorted iteration [2 5 4 6 1 3]
func insertion_sort(values []int, debug bool) []int {
	if !debug {
		for i := 1; i < len(values); i++ {
			key := values[i]
			j := i - 1

			for j >= 0 && values[j] > key {
				values[j+1] = values[j]
				j--
			}
			values[j+1] = key
		}

		return values
	}

	fmt.Printf("BEFORE %v\n", values)

	for i := 1; i < len(values); i++ {
		key := values[i]
		j := i - 1

		fmt.Printf("\n(i := %d | values[i] := %d)", i, values[i])
		fmt.Printf(" key := %d ", key)
		fmt.Printf(" (j := %d | values[j] := %d)\n", j, values[j])

		for j >= 0 && values[j] > key {
			fmt.Printf("doing a swap: values[j+1](%d) := values[j](%d)\n", values[j+1], values[j])

			values[j+1] = values[j]
			j--
			fmt.Printf("swapped iteration %v\n", values)
		}
		fmt.Printf("values[j+1](%d) = key(%d)\n", values[j+1], key)
		values[j+1] = key

		fmt.Printf("sorted iteration %v\n", values)
	}
	fmt.Printf("AFTER %v\n", values)

	return values

}

func insertion_sort_reverse(values []int) []int {
	for i := 1; i < len(values); i++ {
		key := values[i]
		j := i - 1

		for j >= 0 && values[j] < key {
			values[j+1] = values[j]
			j--
		}
		values[j+1] = key
	}

	return values
}

func main() {
	values := []int{5, 2, 4, 6, 1, 3}
	insertion_sort_reverse(values)

	fmt.Printf("values %v\n", values)
}
