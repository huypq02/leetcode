package main

import (
	"fmt"
)

func lemonadeChange(bills []int) bool {
	// changes keeps track of the number of $5 and $10 bills we have
	changes := make(map[int]int, 2)

	// Iterate through each customer's bill
	for _, b := range bills {
		switch b {
		case 5:
			// If customer pays with $5, no change needed, just collect the bill
			changes[5]++
		case 10:
			// If customer pays with $10, need to give $5 as change
			// Greedy: always give the largest bill possible as change
			if changes[5] == 0 {
				// Can't give change
				return false
			}
			changes[5]--
			changes[10]++
		case 20:
			// If customer pays with $20, need to give $15 as change
			// Greedy: prefer to give one $10 and one $5 if possible
			if changes[5] == 0 {
				// Can't give change
				return false
			}
			if changes[10] == 0 {
				// No $10 bill, try to give three $5 bills
				if changes[5] < 3 {
					// Not enough $5 bills
					return false
				}
				changes[5] -= 3
			} else {
				// Give one $10 and one $5 bill
				changes[5]--
				changes[10]--
			}
		}
	}

	// If we never failed to give change, return true
	return true
}

func main() {
	// Example 1
	bills1 := []int{5, 5, 5, 10, 20}
	result1 := lemonadeChange(bills1)
	fmt.Printf("Input: bills = %v\nOutput: %v (Expected: true)\n", bills1, result1)

	// Example 2
	bills2 := []int{5, 5, 10, 10, 20}
	result2 := lemonadeChange(bills2)
	fmt.Printf("Input: bills = %v\nOutput: %v (Expected: false)\n", bills2, result2)

	// Additional test cases
	bills3 := []int{5, 5, 5, 5, 10, 5, 10, 10, 10, 20}
	result3 := lemonadeChange(bills3)
	fmt.Printf("Input: bills = %v\nOutput: %v (Expected: true)\n", bills3, result3)

	bills4 := []int{10, 10}
	result4 := lemonadeChange(bills4)
	fmt.Printf("Input: bills = %v\nOutput: %v (Expected: false)\n", bills4, result4)
}
