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
