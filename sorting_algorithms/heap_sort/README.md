# Heap sort
There are stages to heap sort:
1. Convert input to max heap
2. Extract the elements from the heap one by one and sort the array (this is an in place algorithm, so no extra memory is created, but the sorting process is in place and will modify the input)

## Build the heap
A max heap is a complete binary tree where the parent node is always greater than its children.

example input: `arr = [4, 10, 3, 5, 1]`

### Convert to binary tree ()
The diagram below is the tree representation of the example array, we have not done any modifications or changes with code yet. This is a visual representation of what the array is, and what we will be feeding into the function to max heapify it
```less
       4
      /  \
    10    3
   /  \
  5    1

```

### Heapifying the tree
We start from the last non-leaf node, in code terms we start from the mid-point of the array and descend till we reach the end of the array.

This is because the elements after the mid point are leaf/ child nodes, and everything before the midpoint are parent nodes.
`[4, 10, 3, 5, 1]`

In this case we land at 3, even though I said that every before the midpoint is a parent node, node 3 has the potential to be a parent. Either way, the code will realise that node 3 has no children, so nothing happens and we go down to node 10.

We get this as the max heap, where the parent node is always greater than its children
```less
       10
      /   \
     5     3
   /   \
  4     1

```

### Sorting process (Swap and heapify)
Once the heap is a max heap, we repeatdely remove the root, and place it at the end of the array.

Max heap: `[10 4 3 5 1]`

1. We swap `10` with `1` `[1, 4, 3, 5, 10]`
```less
       1
      /  \
     4    3
    /   
   5

```
1 is not the largest, so heapify and place 4 as the root
```less
       4
      /  \
     1    3
   /   
  5

```

After heapify: `[4, 1, 3, 5, 10]`

### Repeat 
We continue this process of swapping the root going down from the last element from the array, heapifying as we go.
Swap: `[5, 1, 3, 4, 10]`
```less
       5
      /   \
     1     3

```

Heapify: `[3, 1, 5, 4, 10]`

Swap: `[1, 3, 5, 4, 10]`
```less
       3
      /
     1
```
Heapify: `[1, 3, 4, 5, 10]`

Swap: `[3, 1, 4, 5, 10]`
```less
       1
```

Heapify: `[1, 3, 4, 5, 10]`