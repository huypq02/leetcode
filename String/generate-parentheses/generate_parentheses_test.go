package main

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	// List of testcase
	testcases := []struct {
		name           string
		n              int
		expectedResult []string
	}{
		{"n equals 3", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{"n equals 1", 1, []string{"()"}},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := generateParenthesis(tc.n)
			if !reflect.DeepEqual(result, tc.expectedResult) {
				t.Errorf("wrong result %v with n %d; expected %v\n", result, tc.n, tc.expectedResult)
			}
		})
	}
}
