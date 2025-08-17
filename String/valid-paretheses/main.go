package main

type Stack []rune

func (s *Stack) Push(ch rune) {
	*s = append(*s, ch)
}

func (s *Stack) Pop() rune {
	if len(*s) == 0 {
		return 0 // or panic, depending on use case
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func isValid(s string) bool {
	if s[0] == ')' || s[0] == ']' || s[0] == '}' {
		return false
	}

	// Sample the type of the bracket
	var bracket = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := &Stack{}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			(stack).Push(bracket[v])
		}

		if v == ')' || v == ']' || v == '}' {
			if v != (stack).Pop() {
				return false
			}
		}
	}

	return len(*stack) == 0
}

func main() {
	s1 := "()"
	println(isValid(s1)) // Output: true

	s2 := "()[]{}"
	println(isValid(s2)) // Output: true

	s3 := "(]"
	println(isValid(s3)) // Output: false

	s4 := "([])"
	println(isValid(s4)) // Output: true

	s5 := "([)]"
	println(isValid(s5)) // Output: false
}
