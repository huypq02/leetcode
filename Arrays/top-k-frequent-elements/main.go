package main

import (
	"fmt"
	"sort"
)

type element struct {
	value  int
	amount int
}

func topKFrequent(nums []int, k int) []int {

	// Count number of element in the array
	// For example
	// arr = [1, 1, 1, 2, 2, 3]
	// => {1:3, 2:2; 3:1}

	numberOfElement := make(map[int]int)
	for _, num := range nums {
		numberOfElement[num]++
	}

	// Create a slice including the value and the number of it
	var ss []element
	for k, v := range numberOfElement {
		ss = append(ss, element{k, v})
	}

	// Sort the most frequent elements in decreasing order
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].amount > ss[j].amount
	})

	// Filter the k most frequent elements
	var topKFrequent []int
	for i := 0; i < k; i++ {
		topKFrequent = append(topKFrequent, ss[i].value)
	}

	return topKFrequent
}

func main() {
	var nums, k = []int{1, 1, 1, 2, 2, 3}, 2
	fmt.Println(topKFrequent(nums, k))

	var nums2, k2 = []int{1}, 1
	fmt.Println(topKFrequent(nums2, k2))
}
