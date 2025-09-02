package main

import "fmt"

func countGoodSubstrings(s string) int {
	fixed := 3
	count := 0
	for right := fixed; right <= len(s); right++ {
		if s[right-1] != s[right-2] && s[right-2] != s[right-3] && s[right-1] != s[right-3] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countGoodSubstrings("xyzzaz"))
	fmt.Println(countGoodSubstrings("aababcabc"))
	fmt.Println(countGoodSubstrings("aaababcabc"))
	fmt.Println(countGoodSubstrings("aaababcabc"))
	fmt.Println(countGoodSubstrings("aaababcabc"))
	fmt.Println(countGoodSubstrings("aaababcabc"))
	fmt.Println(countGoodSubstrings("aaababcabc"))
	fmt.Println(countGoodSubstrings("aaababcabc"))
	fmt.Println(countGoodSubstrings("aaababcabc"))
}
