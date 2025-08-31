package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	left, maxProfit := 0, 0
	for right := 1; right < len(prices); right++ {
		if prices[left] > prices[right] {
			left = right
			continue
		}
		profit := prices[right] - prices[left]
		if maxProfit < profit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func main() {
	// Input: prices = [7,1,5,3,6,4]
	prices := []int{7, 1, 5, 3, 6, 4}
	// Output: 5
	println(maxProfit(prices))

	// Input: prices = [7,6,4,3,1]
	prices2 := []int{7, 6, 4, 3, 1}
	// Output: 0
	println(maxProfit(prices2))

	// Input: [2,4,1]
	prices3 := []int{2, 4, 1}
	// Output: 2
	println(maxProfit(prices3))
}
