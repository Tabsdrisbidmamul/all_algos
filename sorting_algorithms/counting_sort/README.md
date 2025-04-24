# Counting Sort

A very simple algorithm, different from the comparision algorithms like merge, quick, and heap sort. This is an `O(n)` algorithm, and relies on the ability that are duplicate elements in the array.

## Finding the max element phase
We take in array, and the first element is the `max` element. We then loop through the array, to find the real `max` element.
Once done, create a counting array that is `max + 1`.

## Creating the count array phase
We then loop through the original array, and place each element found as `count[num]++`. The index of the count is the original element, and it stores a running count of elements found.

## Sorting the original element
- We initialise a `sorted_index` count to 0, this is the index that we use to move the elements into their sorted order.
- We then do a double loop, one over the counting array, and a while loop till `frequency > 0`.
- We do `arr[sorted_index] = i`, where `i` is the index of the `count` array. Decrement `frequency` variable, and increment `sorted_index`

## Final Result
The original array has been sorted, according to the steps above.

Issue is that this algorithm creates another array of `max + 1` i.e the maximum value in the original array could be 100, so we have a counting array that is of size 101, and could very well have empty positions because counting sort will count on the index.
