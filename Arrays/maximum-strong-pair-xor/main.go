package main

import (
	"fmt"
)

func maximumStrongPairXor(nums []int) int {
	maxXor := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			x, y := nums[i], nums[j]
			if abs(x-y) <= min(x, y) {
				if x^y > maxXor {
					maxXor = x ^ y
				}
			}
		}
	}

	return maxXor
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(maximumStrongPairXor(nums))

	nums = []int{10, 100}
	fmt.Println(maximumStrongPairXor(nums))

	nums = []int{5, 6, 25, 30}
	fmt.Println(maximumStrongPairXor(nums))

	nums = []int{1, 1, 3, 8, 7}
	fmt.Println(maximumStrongPairXor(nums))

	nums = []int{1, 1, 10, 3, 9}
	fmt.Println(maximumStrongPairXor(nums))
}
