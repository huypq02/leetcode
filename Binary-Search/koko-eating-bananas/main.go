package main

func minEatingSpeed(piles []int, h int) int {
	// Length of the array
	n := len(piles)

	// Find the max element of the array
	max := piles[0]
	for i := 1; i < n; i++ {
		if max < piles[i] {
			max = piles[i]
		}
	}

	left, right, k := 1, max, max
	for left <= right {
		mid := (right-left)/2 + left // Binary search
		hoursNeeded := 0
		for _, v := range piles { // Check each element of the piles
			hoursNeeded += v / mid
			if v%mid != 0 {
				hoursNeeded++
			}
		}

		if hoursNeeded <= h { // Found a valid hour, try to find a smaller k
			right = mid - 1
			if k > mid {
				k = mid // Save the minimum integer k such that she can eat all the bananas within h hours
			}
		} else if hoursNeeded > h { // Invalid hour
			left = mid + 1
		}
	}
	return k
}

func main() {
	// Example 1
	piles1 := []int{3, 6, 7, 11}
	h1 := 8
	result1 := minEatingSpeed(piles1, h1)
	println("Example 1 Output:", result1, "expected: 4")

	// Example 2
	piles2 := []int{30, 11, 23, 4, 20}
	h2 := 5
	result2 := minEatingSpeed(piles2, h2)
	println("Example 2 Output:", result2, "expected: 30")

	// Example 3
	piles3 := []int{30, 11, 23, 4, 20}
	h3 := 6
	result3 := minEatingSpeed(piles3, h3)
	println("Example 3 Output:", result3, "expected: 23")

	// Example 4
	piles4 := []int{332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184}
	h4 := 823855818
	result4 := minEatingSpeed(piles4, h4)
	println("Example 4 Output:", result4, "expected: 14")
}
