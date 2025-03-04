package main

import "fmt"

func checkValid(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

		}
	}

	return true
}

func main() {
	m1 := [][]int{
		{1, 2, 3},
		{3, 1, 2},
		{2, 3, 1},
	}
	fmt.Println(checkValid(m1))

	m2 := [][]int{
		{1, 1, 1},
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println(checkValid(m2))
}
