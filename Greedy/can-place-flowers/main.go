package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	// Greedy approach: try to plant a flower at every possible spot
	// by always making the locally optimal choice (planting whenever possible)
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		// If the left neighbor is planted, skip this spot
		if i > 0 && flowerbed[i-1] == 1 {
			continue
		}
		// If the right neighbor is planted, skip this spot
		if i < len(flowerbed)-1 && flowerbed[i+1] == 1 {
			continue
		}

		// If current spot is empty and both neighbors are not planted,
		// plant a flower here (count it)
		if flowerbed[i] == 0 {
			count++
		}

		// Move to the next non-adjacent spot to avoid planting flowers next to each other
		i++
	}
	// If we can plant at least n flowers, return true
	return count >= n
}

func main() {
	// Example 1
	flowerbed1 := []int{1, 0, 0, 0, 1}
	n1 := 1
	// Expected: true
	println("Testcase 1:", canPlaceFlowers(flowerbed1, n1), "Expected: true")

	// Example 2
	flowerbed2 := []int{1, 0, 0, 0, 1}
	n2 := 2
	// Expected: false
	println("Testcase 2:", canPlaceFlowers(flowerbed2, n2), "Expected: false")

	// Additional testcases
	flowerbed3 := []int{0, 0, 1, 0, 0}
	n3 := 2
	// Expected: true
	println("Testcase 3:", canPlaceFlowers(flowerbed3, n3), "Expected: true")

	flowerbed4 := []int{0}
	n4 := 1
	// Expected: true
	println("Testcase 4:", canPlaceFlowers(flowerbed4, n4), "Expected: true")

	flowerbed5 := []int{1}
	n5 := 1
	// Expected: false
	println("Testcase 5:", canPlaceFlowers(flowerbed5, n5), "Expected: false")

	// Additional testcase
	flowerbed6 := []int{1, 0, 0, 0, 0, 1}
	n6 := 2
	// Expected: false
	println("Testcase 6:", canPlaceFlowers(flowerbed6, n6), "Expected: false")
}
