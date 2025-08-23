package main

import (
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
	var smallestString, longestString string
	if len(str1) < len(str2) {
		smallestString, longestString = str1, str2
	} else {
		smallestString, longestString = str2, str1
	}

	var gcdOfStrings, nGcdOfStrings string
	for i := len(smallestString); i > 0 && len(nGcdOfStrings) != len(smallestString); i-- {
		t := smallestString[:i]
		if len(smallestString)%len(t) != 0 {
			continue // because t is concatenated with itself N-times
		}

		if len(t) == 0 {
			continue
		}
		// Concatenate itself
		tmp := strings.Repeat(t, len(longestString)/len(t))
		if longestString == tmp {
			gcdOfStrings = t
			break
		}
	}

	// Re-check the smallest string
	if len(gcdOfStrings) == 0 {
		return ""
	}
	nGcdOfStrings = strings.Repeat(gcdOfStrings, len(smallestString)/len(gcdOfStrings))
	if nGcdOfStrings != smallestString {
		return ""
	}
	return gcdOfStrings
}

func main() {
	str1a, str2a := "ABCABC", "ABC"
	println(gcdOfStrings(str1a, str2a)) // Output: "ABC"

	str1b, str2b := "ABABAB", "ABAB"
	println(gcdOfStrings(str1b, str2b)) // Output: "AB"

	str1c, str2c := "LEET", "CODE"
	println(gcdOfStrings(str1c, str2c)) // Output: ""

	str1d, str2d := "ABABABAB", "ABAB"
	println(gcdOfStrings(str1d, str2d)) // Output: "ABAB"

	str1e, str2e := "TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"
	println(gcdOfStrings(str1e, str2e)) // Output: "TAUXX"

	str1f, str2f := "ABCABCABC", "ABCAAA"
	println(gcdOfStrings(str1f, str2f)) // Output: ""
}
