# Radix Sort

We are sorting numbers by units so:
- 1s
- 10s
- 100s

## Get max element
Iterate through the whole array to find the maximum element

## For loop over counting sort
We have the max element, so we do a for loop
we start from the 1s position and move up to 10s, then 100s, etc.

```go
for exp := 1; max/exp > 0; exp *= 10 {
		arr = counting_sort(arr, exp)
}
```

Each time we are doing a stable sort on the original array, moving the elements to the correct position

## stable sort
We create 2 arrays:
- count array
  - This is an array from `0-9` so the digits each number ends in
- output array: The sorted array
  
### count digits
We first do a loop on the input array, and get the digit postion depending on the unit we are in the loop.

So if doing the loop for 1s, we do
```go
digit := (arr[i] / exp) % 10
```
and add the digit's position to the count array, remember the index is the digit and the count array keeps track of the number of occurrences.


### Convert frequncy to positions 
We do a loop over the count array to convert the frequencies to positions.
```go
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
```

example input: `[12, 10, 15, 21]`
`count = [1 1 1 0 0 1 0 0 0 0]`

Converting to positions, we get this array
```less
count = [1 2 3 3 3 4 4 4 4 4]
         ↑ ↑ ↑     ↑
         0 1 2     5
```

What this means:
- Numbers ending in 0 go to position 0
- Numbers ending in 1 go to position 1
- Numbers ending in 2 go to position 2
- Numbers ending in 5 go to position 3

See here that `3` repeated 3 times, this tells us that 2 things:
1. any numbers ending in 2 always go to position 2
2. This happens 3 times, so we either move numbers or do nothing

### Doing the move
```go
	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / exp) % 10
		output[count[digit]-1] = arr[i]
		count[digit]--
	}
```

We have the positions from above, so we do iteration from the end of the original array.

The index will always be inbound of the count array, since we are doing a modular by `10`