package main

func searchMatrix(matrix [][]int, target int) bool {
	rowIdx := 0
	// Find the row that may contain the target
	for i := 0; i < len(matrix)-1; i++ {
		if matrix[i][0] < target && matrix[i+1][0] > target {
			rowIdx = i
		} else if matrix[i+1][0] <= target {
			rowIdx = i + 1
		}
	}

	// Search for the target in the identified row
	for j := 0; j < len(matrix[rowIdx]); j++ {
		if matrix[rowIdx][j] == target {
			return true
		}
	}

	return false
}

func main() {
	// Example 1
	matrix1 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target1 := 3
	result1 := searchMatrix(matrix1, target1)
	println("Example 1 Output:", result1) // Expected: true

	// Example 2
	matrix2 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target2 := 13
	result2 := searchMatrix(matrix2, target2)
	println("Example 2 Output:", result2) // Expected: false

	// Example 3
	matrix3 := [][]int{{1}, {3}}
	target3 := 3
	result3 := searchMatrix(matrix3, target3)
	println("Example 3 Output:", result3) // Expected: true

	// Example 4
	matrix4 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	target4 := 30
	result4 := searchMatrix(matrix4, target4)
	println("Example 4 Output:", result4) // Expected: true
}
