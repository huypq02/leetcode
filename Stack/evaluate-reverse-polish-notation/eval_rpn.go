package main

import "strconv"

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val) // Append a new value to the slice (dynamic array)
}

func (s *Stack) Pop() int {
	top := (*s)[len(*s)-1] // Top is the last element of the slice
	*s = (*s)[:len(*s)-1]  // Truncate the slice by reducing its length by one element
	return top
}

func evalRPN(tokens []string) int {
	// Here is an approach to solve the exercise
	//
	// For instance:
	// ["2","1","+","3","*"]
	// if it's a stack
	// 2 => Push => 1 => Push => + => Pop twice (1 and 2): 2 + 1 = 3 => Push
	// 3 => Push => * => Pop twice (3 and 3): 3 * 3 = 3

	// For another example:
	// ["4","13","5","/","+"]
	// 4 => Push =>  13 => Push => 5 Push => / => Pop twice (5 and 13): 13 / 5 = 2 => Push
	// + => Pop twice (2 and 4) => 4 + 2 = 6

	// Another one
	// ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
	// 10 => Push => 6 => Push => 9 => Push => 3 => Push => + => Pop twice (3 and 9): 9 + 3 = 12 => Push
	// -11 => Push => * => Pop twice (-11 and 12): 12 * -11 = -132 => Push
	// / => Pop twice (-132 and 6):  6 /-132 = 0 => Push
	// * => Pop twice (0 and 10):  10 *0 = 0 => Push
	// 17 => Push => + => Pop twice (17 and 0): 0 + 17  = 17 => Push
	// 5 => Push => + => Pop twice (5 and 17): 17 * 5  = 22

	if len(tokens) == 1 {
		num, err := strconv.Atoi(tokens[0])
		if err != nil {
			return 0 // something wrong
		}

		return num
	}

	// Define a stack
	stack := Stack{}
	result := 0
	// Retrieving each token in the array
	for _, v := range tokens {
		// Switch case each operator and numbers
		switch v {
		case "+": // Plus
			num2 := stack.Pop()
			num1 := stack.Pop()
			result = num1 + num2
			stack.Push(result)
		case "-": // Minus
			num2 := stack.Pop()
			num1 := stack.Pop()
			result = num1 - num2
			stack.Push(result)
		case "*": // Multiply
			num2 := stack.Pop()
			num1 := stack.Pop()
			result = num1 * num2
			stack.Push(result)
		case "/": // Divide
			num2 := stack.Pop()
			num1 := stack.Pop()
			result = num1 / num2
			stack.Push(result)
		default: // not be operator
			num, err := strconv.Atoi(v)
			if err != nil {
				return 0 // something wrong
			}
			stack.Push(num)
		}
	}

	return result
}
