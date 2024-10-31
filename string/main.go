package main

import (
	"fmt"
)

func commonPrefix(str1, str2 string) string {
	minLen := len(str1)
	if len(str2) < minLen {
		minLen = len(str2)
	}

	for i := 0; i < minLen; i++ {
		if str1[i] != str2[i] {
			return str1[:i]
		}
	}
	return str1[:minLen]
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		prefix = commonPrefix(prefix, strs[i])
	}

	return prefix
}

func main() {

	str := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(str))

	str1 := []string{"", "b"}
	fmt.Println(longestCommonPrefix(str1))

	str2 := []string{"a", "b"}
	fmt.Println(longestCommonPrefix(str2))

	str3 := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(str3))
}
