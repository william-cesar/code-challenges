### Rotate Image

Note: Try to solve this task in-place (with O(1) additional memory), since this is what you'll be asked to do during an interview.

You are given an n x n 2D matrix that represents an image. Rotate the image by 90 degrees (clockwise).

> Example

- For

```
a = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9]
]
```

the output should be

```
solution(a) = [
  [7, 4, 1],
  [8, 5, 2],
  [9, 6, 3]
]
```

> Input/Output

[execution time limit] 4 seconds (go)

[memory limit] 1 GB

[input] array.array.integer a

Guaranteed constraints:

- 1 ≤ a.length ≤ 100,
- a[i].length = a.length,
- 1 ≤ a[i][j] ≤ 10^4.

[output] array.array.integer

### STC

> Space: O(1)

> Time: O(n²)
