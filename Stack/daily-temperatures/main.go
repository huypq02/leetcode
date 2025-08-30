package main

import "fmt"

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val) // Append a new value to the slice (dynamic array)
}

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1] // Truncate the slice by reducing its length by one element
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1] // Top is the last element of the slice
}

func dailyTemperatures(temperatures []int) []int {
	// array [73, 74, 75, 71, 69, 72, 76, 73]
	// dailyTemperatures [0, 0, 0, 0, 0, 0, 0, 0]
	// i=0 | stack = [0]

	// i=1 | loop | prevIndex = 0
	//            | stack = [] // pop
	//            | dailyTemperatures [(1-0), 0, 0, 0, 0, 0, 0, 0]
	//     | stack = [1]
	//
	// i=2 | loop | prevIndex = 1
	//            | stack = [] // pop
	//            | dailyTemperatures [1, (2-1), 0, 0, 0, 0, 0, 0]
	//     | stack = [2]
	//
	// i=3 | stack = [2,3]
	//
	// i=4 | stack = [2,3,4]
	//
	// i=5 | loop1| prevIndex = 4
	//            | stack = [2,3] // pop
	//            | dailyTemperatures [1, 1, 0, 0, 5-4, 0, 0, 0]
	//       loop2| prevIndex = 3
	//            | stack = [2] // pop
	//            | dailyTemperatures [1, 1, 0, 5-3, 1, 0, 0, 0]
	//     | stack = [2, 5]
	//
	// i=6 | loop1| prevIndex = 5
	//            | stack = [2] // pop
	//            | dailyTemperatures [1, 1, 0, 2, 1, 6-5, 0, 0]
	//     | loop2| prevIndex = 2
	//            | stack = [] //pop
	//            | dailyTemperatures [1, 1, 6-2, 2, 1, 1, 0, 0]
	// i=7 | stack = [7]
	//
	// return dailyTemperatures is [1, 1, 6-2, 2, 1, 1, 0, 0]

	n := len(temperatures)
	// Define a stack for the daily temperature
	dailyTemperatures := make([]int, n)
	// Stack to store previous index
	stack := Stack{}

	// Retrieving an array of temperatures
	for i := 0; i < n; i++ {
		// Loop until finding the day to get a warmer temperature
		for len(stack) > 0 && temperatures[i] > temperatures[stack.Top()] {
			prevIndex := stack.Top()
			stack.Pop()                                  // pop
			dailyTemperatures[prevIndex] = i - prevIndex // the day with prevIndex stored before
		}
		stack.Push(i)
	}

	return dailyTemperatures
}

func main() {
	temperatures1 := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures1)) // Output [1,1,4,2,1,1,0,0]

	temperatures2 := []int{30, 40, 50, 60}
	fmt.Println(dailyTemperatures(temperatures2)) // Output [1,1,1,0]

	temperatures3 := []int{30, 60, 90}
	fmt.Println(dailyTemperatures(temperatures3)) // Output [1,1,0]
}
