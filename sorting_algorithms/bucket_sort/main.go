package main

import (
	"fmt"
	"sort"
)

func bucket_sort(arr []float64) []float64 {
	n := len(arr)
	if n <= 0 {
		return arr
	}

	// make n buckets (i.e length of array is 10 so 10 buckets)
	buckets := make([][]float64, n)

	// place items into either of the 10 buckets, so 0.11, 0.13, 0.15 all go into bucket 1 (index 1)
	// items 0.65, 0.62 go into bucket 6 (index 6) and so
	for _, num := range arr {
		index := int(num * float64(n))
		buckets[index] = append(buckets[index], num)
	}

	result := []float64{}
	for _, bucket := range buckets {
		sort.Float64s(bucket)
		result = append(result, bucket...)
	}

	return result

}

func main() {
	arr := []float64{0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68}
	fmt.Println("unsorted array:", arr)
	sorted := bucket_sort(arr)
	fmt.Println("sorted array:", sorted)
}
