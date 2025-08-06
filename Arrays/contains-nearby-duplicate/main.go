package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[v]; ok && abs(i-j) <= k {
			fmt.Print(m)
			fmt.Print("\n")
			fmt.Print(i)
			fmt.Print("\n")
			fmt.Print(j)
			fmt.Print("\n")
			return true
		}
		m[v] = i
	}

	return false
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num

}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3)) // true

	// m[1] = 0
	// m[2] = 1
	// m[3] = 2
	// m[1] = 3
	// 3 - 0 = 3 <= 1
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))       // true
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)) // false
}
