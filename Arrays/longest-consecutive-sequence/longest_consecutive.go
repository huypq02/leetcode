package main

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}
	// Sort the array
	// Ex: [100,4,200,1,3,2]
	sortedNums := sort.IntSlice(nums)
	sortedNums.Sort()

	// Handle with the stored array of the numbers
	// Ex: [1 2 3 4 100 200]
	countLongestConsecutive := 1
	tempCount := 1
	for i := 0; i < len(sortedNums)-1; i++ {
		if sortedNums[i] == sortedNums[i+1] {
			continue
		}
		if sortedNums[i+1]-sortedNums[i] > 1 {
			if countLongestConsecutive < tempCount {
				countLongestConsecutive = tempCount
			}
			tempCount = 1 // reset
			continue
		}
		tempCount++
	}

	if countLongestConsecutive < tempCount {
		return tempCount
	}
	return countLongestConsecutive
}
