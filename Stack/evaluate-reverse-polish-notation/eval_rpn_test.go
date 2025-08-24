package main

import "testing"

func Test_evalRPN(t *testing.T) {
	testcases := []struct {
		name           string
		tokens         []string
		expectedResult int
	}{
		{"plus and multiply", []string{"2", "1", "+", "3", "*"}, ((2 + 1) * 3)},
		{"divide and plus", []string{"4", "13", "5", "/", "+"}, (4 + (13 / 5))},
		{"all operators",
			[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			((10 * (6 / ((9 + 3) * -11))) + 17) + 5},
		{"have not any operators", []string{"18"}, 18},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := evalRPN(tc.tokens)

			if result != tc.expectedResult {
				t.Errorf("wrong result %d with tokens %v, expected %d",
					result,
					tc.tokens,
					tc.expectedResult)
			}
		})
	}
}
