package main

import "fmt"

func productExceptSelf(nums []int) []int {
	productOfAllExceptZero := 1
	productExceptSelf := make([]int, len(nums))
	countZero := 0
	for _, v := range nums {
		if v == 0 {
			countZero++
			continue
		}
		productOfAllExceptZero *= v
	}

	switch countZero {
	case 0:
		productExceptSelf = nil
		for _, v := range nums {
			productExceptSelf = append(productExceptSelf, productOfAllExceptZero/v)
		}
	case 1:
		for k, v := range nums {
			if v != 0 {
				continue
			}
			productExceptSelf[k] = productOfAllExceptZero
		}
	}

	return productExceptSelf
}

func main() {
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums1)) // Output: [24,12,8,6]

	nums2 := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums2)) // Output: [0,0,9,0,0]
}
