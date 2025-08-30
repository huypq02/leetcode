package main

func backtrack(path []rune, remainingLeft, remainingRight int, result *[]string) {

	// Validate the wrong paths that isn't proper
	if remainingLeft < 0 || remainingRight < 0 || remainingLeft > remainingRight {
		return
	}

	// Add the proper paths to the result
	if remainingLeft+remainingRight == 0 {
		*result = append(*result, string(path))
		return
	}

	// Define a list of a parenthesis
	parenthesis := []rune{'(', ')'}
	// Retrieving the list
	for _, v := range parenthesis {
		// Add to the path
		path = append(path, v)
		switch v {
		case '(':
			// Explore the next path with the parenthesis '('
			backtrack(path, remainingLeft-1, remainingRight, result)
		case ')':
			// Explore the next path with the parenthesis ')'
			backtrack(path, remainingLeft, remainingRight-1, result)
		}

		// Un-choose
		path = path[:len(path)-1]
	}
}

func generateParenthesis(n int) []string {
	// Declare a list of the results
	res := make([]string, 0)

	backtrack([]rune{}, n, n, &res)
	return res
}
