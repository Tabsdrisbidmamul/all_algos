# Loop Invariants
A loop invariant is a condition or property that holds true before and after each iteration of a loop in an algorithm. It helps ensure the correctness of the algorithm. By defining a loop invariant, you can reason about the behavior of the loop and how the algorithm progresses toward its goal.

## Key Characteristics of Loop Invariants:

- Initialization: The invariant must hold true before the loop starts (at the very beginning).
- Maintenance: It must remain true after each iteration of the loop.
- Termination: When the loop ends, the invariant can help prove that the algorithm produces the correct result.

Loop invariants are often used to prove the correctness of algorithms, especially when designing or analyzing algorithms that involve loops.

## Insertion Sort Loop Invaraints
example: [5, 2, 4, 6, 1, 3]
- Initialisation:
  - We start the outer loop from 1..arr.length, and the inner loop is i-1..=i. It is always true that our inner loop will contain a sorted sub-array
  - The first iteration, i = 1, and j = 0, and for example above, 5 is sorted, since one element sub array is already sorted.
- Maintanance:
  - The inner loop will start shifting values within the sub array. 
  - When (i := 1 | values[i] := 2) key := 2  (j := 0 | values[j] := 5)
  the swapped iteration looks like this swapped iteration [5 5 4 6 1 3], and when the inner loop terminates and before the outer loop terminates we will do one last swap to get the sorted iteration [2 5 4 6 1 3]
  - The sub array will always be sorted see that [2, 5] are in the correct order.
- Termination:
  - Once the outer loop finishes completely we get the complete sorted array [1 2 3 4 5 6], such that the sub-array has moved all the elements to their correct order so the full array is in the correct order.