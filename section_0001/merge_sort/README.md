# Divide and conquer
The basic idea is broken into 3 problems
- Divide: 
  - divide the problem into sub-problems that are smaller instances of the larger problem
- Conquer:
  - solve the sub-problems recursively
- Combine:
  - combine the solved sub-problems to form a solution to the original problem

## merge sort
Will divide the array into halfs, until an array contains one element. We recursively combine each divided array, sorting as we go. Such that each combined sub-array's left and right will always be sorted correctly when combining.

### 1. Divide
Divide the array up till there is only one element (base case), and then move onto phase 2 (conquer)
Example Input Array:
`[38, 27, 43, 3, 9, 82, 10]`
```less
[38, 27, 43, 3, 9, 82, 10]
         /               \
[38, 27, 43, 3]      [9, 82, 10]
      /     \          /     \
[38, 27]  [43, 3]   [9, 82]  [10]
  /   \    /   \    /   \
[38]  [27] [43] [3] [9]  [82]  [10]
 ```

 ### 2. Conquer
 Loop and compare each element in 2 arrays (i, j -> counters), and place each finding into a results array 
 ```less
[38]  [27] → [27, 38]
[43]  [3]  → [3, 43]
[9]   [82] → [9, 82]
[10]         → [10] (unchanged)
 ```

### 3. Combine
Combine each of the arrays (this is handled by the recusion in the algorithm)
```less
[27, 38]  [3, 43] → [3, 27, 38, 43]
[9, 82]   [10]    → [9, 10, 82]
```

```less
[3, 27, 38, 43]  [9, 10, 82] → [3, 9, 10, 27, 38, 43, 82]
```

```less
[3, 9, 10, 27, 38, 43, 82]
```

## Time Complexity
At each level, the array is split into two halves, and merging takes `O(n)` time.
Since we do `log(n)` levels of splitting, the overall time complexity is:

`O(nlogn)`

Space Complexity: `O(n)` (for the temporary arrays during merging).