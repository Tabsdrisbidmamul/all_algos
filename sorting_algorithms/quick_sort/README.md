# Quick Sort
A divide and conquer sorting algorithm, it works by selecting a pivot point. Everything below the pivot is less than it and everything above the pivot is greater than it. 

It is recursively applied to the subarrays until it is sorted.

## Choose a pivot
On the first run this is the last element
`[8, 4, 7, 2, 9, 5, 4]`

`pivot=3`

## Parition the array
We move elements smaller than the pivot to the left and elements greater to the right.
```less
Before Partitioning:
[5, 3, 8, 2, 7]
                             
Partitioning around pivot: 7
[5, 3, 2, 7, 8] 
```

### Partition (1)
- Starting array: `[1, 5, 2, 6]`
#### 1st loop
- `low = 0 -> arr[low] = 1` `high = 3 -> arr[high] = 6`
- `pivot = 6`, `i = 0`
- loop `j := low; j < high; j++ -> j = 0`
- `1 < 6`: swap `arr[i], arr[j] -> arr[0], arr[0]`. `i++ -> i = 1`

#### 2nd loop
- `low = 0 -> arr[low] = 1` `high = 3 -> arr[high] = 6`
- `pivot = 6`, `i = 1`
- loop `j := low; j < high; j++ -> j =1`
- `5 < 6`: swap `arr[i], arr[j] -> arr[1], arr[1]`. `i++ -> i = 2`

#### 3rd loop
- `low = 0 -> arr[low] = 1` `high = 3 -> arr[high] = 6`
- `pivot = 6`, `i = 2`
- loop `j := low; j < high; j++ -> j = 2`
- `2 < 6`: swap `arr[i], arr[j] -> arr[2], arr[2]`. `i++ -> i = 3`

#### loop termination
- `low = 0 -> arr[low] = 1` `high = 3 -> arr[high] = 6`
- `pivot = 6`, `i = 3`
- swap `arr[i], arr[high] = arr[3], arr[3]`

#### final array and new pivot
- Array: `[1, 5, 2, 6]`
- It is correct since all elements below the pivot `6` are less than it
- `i = 3`

### Sorting left sub array (1)
- Starting array: `[1, 5, 2, 6]`
#### 1st loop
- `low = 0 -> arr[low] = 1` `high = 2 -> arr[high] = 2`
- `pivot = 2`, `i = 0`
- loop `j := low; j < high; j++ -> j = 0`
- `1 < 2`: swap `arr[i], arr[j] -> arr[0], arr[0]`. `i++ -> i = 1`

#### 2nd loop
- `low = 0 -> arr[low] = 1` `high = 2 -> arr[high] = 2`
- `pivot = 2`, `i = 1`
- loop `j := low; j < high; j++ -> j = 1`
- `5 < 2`: false

#### loop termination
- `low = 0 -> arr[low] = 1` `high = 2 -> arr[high] = 2`
- `pivot = 2`, `i = 1`
- swap `arr[i], arr[high] = arr[1], arr[2]`

#### final array and new pivot
- Array: `[1, 2, 5, 6]`
- It is correct since all elements below the pivot `2` are less than it and everything above it is greater than the pivot
- `i = 1`

#### Sorting right sub array (1)
In the code, since it is a recursive function, when we get back from `quick_sort(arr, 0, p(3)-1)` the stack frame for sorting the right sub array. `p = 3`, and we pass in `quick_sort(arr, p(3)+1, 3)`.

The if check will stop it from continuing to sort, because the array has already been sorted.
