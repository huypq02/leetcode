# Search a 2D Matrix

**Problem** [LeetCode 74 - Search a 2D Matrix](https://leetcode.com/problems/search-a-2d-matrix/description/)

**Difficulty:** Medium

---

You are given an m x n integer matrix `matrix` with the following two properties:

- Each row is sorted in non-decreasing order.
- The first integer of each row is greater than the last integer of the previous row.

Given an integer `target`, return `true` if `target` is in `matrix` or `false` otherwise.

You must write a solution in $O(\log(m \times n))$ time complexity.

---

## Example 1

**Input:**

```
matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]]
target = 3
```

**Output:**

```
true
```

## Example 2

**Input:**

```
matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]]
target = 13
```

**Output:**

```
false
```

---

## Constraints

- $m = \text{matrix.length}$
- $n = \text{matrix[i].length}$
- $1 \leq m, n \leq 100$
- $-10^4 \leq \text{matrix[i][j]}, \text{target} \leq 10^4$
