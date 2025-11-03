package main

import "fmt"

func isLongPressedName(name string, typed string) bool {
	// Two pointers: namePtr for 'name', typedPtr for 'typed'
	namePtr, typedPtr := 0, 0
	// Iterate through 'typed' string
	for typedPtr < len(typed) {
		// If current characters match, move both pointers forward
		if namePtr < len(name) && name[namePtr] == typed[typedPtr] {
			typedPtr++
			namePtr++
			// If current 'typed' character is a long press (same as previous 'typed'), move typedPtr forward
		} else if typedPtr > 0 && typed[typedPtr] == typed[typedPtr-1] {
			typedPtr++
			// If neither condition is met, it's not a valid long press scenario
		} else {
			return false
		}
	}

	// All characters in name must be matched
	return namePtr == len(name)
}

func main() {
	// Example 1
	name1 := "alex"
	typed1 := "aaleex"
	fmt.Printf("Test 1 - Input: name=\"%s\", typed=\"%s\"\n", name1, typed1)
	fmt.Printf("Expected: true, Actual: %v\n\n", isLongPressedName(name1, typed1))

	// Example 2
	name2 := "saeed"
	typed2 := "ssaaedd"
	fmt.Printf("Test 2 - Input: name=\"%s\", typed=\"%s\"\n", name2, typed2)
	fmt.Printf("Expected: false, Actual: %v\n\n", isLongPressedName(name2, typed2))

	// Example 3
	name3 := "leelee"
	typed3 := "lleeelee"
	fmt.Printf("Test 3 - Input: name=\"%s\", typed=\"%s\"\n", name3, typed3)
	fmt.Printf("Expected: true, Actual: %v\n\n", isLongPressedName(name3, typed3))

	// Example 4
	name4 := "alex"
	typed4 := "aaleexa"
	fmt.Printf("Test 4 - Input: name=\"%s\", typed=\"%s\"\n", name4, typed4)
	fmt.Printf("Expected: false, Actual: %v\n\n", isLongPressedName(name4, typed4))

	// Example 5
	name5 := "alexd"
	typed5 := "ale"
	fmt.Printf("Test 5 - Input: name=\"%s\", typed=\"%s\"\n", name5, typed5)
	fmt.Printf("Expected: false, Actual: %v\n\n", isLongPressedName(name5, typed5))
}
