package main

import "fmt"

func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	max := 0
	for k, v := range m {
		if _, ok := m[k+1]; ok {
			if v+m[k+1] > max {
				max = v + m[k+1]
			}
		}
	}

	return max
}

func main() {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	fmt.Println(findLHS(nums)) // 5

	nums = []int{1, 2, 3, 4}
	fmt.Println(findLHS(nums)) // 2

	nums = []int{1, 1, 1, 1}
	fmt.Println(findLHS(nums)) // 0
}
