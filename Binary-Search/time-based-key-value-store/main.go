package main

import "fmt"

type entry struct {
	value     string
	timestamp int
}

type TimeMap struct {
	m map[string][]entry
}

func Constructor() TimeMap {
	return TimeMap{m: make(map[string][]entry)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	e := entry{value, timestamp}
	this.m[key] = append(this.m[key], e)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	entries, exists := this.m[key] // Check the key exists or not
	if !exists {
		return ""
	}

	result := ""
	// Using Binary search because all the timestamps of set are strictly increasing
	left, right := 0, len(entries)-1 // Set the left and right pointers
	for left <= right {
		mid := left + (right-left)/2            // Set the middle pointer to binary search
		timestampPrev := entries[mid].timestamp // Assign a value such that set was called previously

		if timestampPrev <= timestamp {
			left = mid + 1
			result = entries[mid].value
		} else {
			right = mid - 1
		}
	}
	return result
}

func main() {
	// Test case from README.md example
	fmt.Println("Running test case from README.md:")
	fmt.Println("Input: [\"TimeMap\", \"set\", \"get\", \"get\", \"set\", \"get\", \"get\"]")
	fmt.Println("       [[], [\"foo\", \"bar\", 1], [\"foo\", 1], [\"foo\", 3], [\"foo\", \"bar2\", 4], [\"foo\", 4], [\"foo\", 5]]")

	// Initialize TimeMap
	timeMap := Constructor()
	fmt.Print("Output: [")
	if len(timeMap.m) == 0 {
		fmt.Printf("null")
	}

	// timeMap.set("foo", "bar", 1)
	timeMap.Set("foo", "bar", 1)
	fmt.Print(", null")

	// timeMap.get("foo", 1) -> should return "bar"
	result1 := timeMap.Get("foo", 1)
	fmt.Printf(", \"%s\"", result1)

	// timeMap.get("foo", 3) -> should return "bar"
	result2 := timeMap.Get("foo", 3)
	fmt.Printf(", \"%s\"", result2)

	// timeMap.set("foo", "bar2", 4)
	timeMap.Set("foo", "bar2", 4)
	fmt.Print(", null")

	// timeMap.get("foo", 4) -> should return "bar2"
	result3 := timeMap.Get("foo", 4)
	fmt.Printf(", \"%s\"", result3)

	// timeMap.get("foo", 5) -> should return "bar2"
	result4 := timeMap.Get("foo", 5)
	fmt.Printf(", \"%s\"", result4)

	fmt.Println("]")

	// Expected output: [null, null, "bar", "bar", null, "bar2", "bar2"]
	fmt.Println("Expected: [null, null, \"bar\", \"bar\", null, \"bar2\", \"bar2\"]")

	// Additional test cases
	fmt.Println("\nAdditional test cases:")

	// Test case 2: Empty key lookup
	timeMap2 := Constructor()
	empty := timeMap2.Get("nonexistent", 1)
	fmt.Printf("Get non-existent key: \"%s\" (should be empty)\n", empty)

	// Test case 3: Multiple keys
	timeMap3 := Constructor()
	timeMap3.Set("key1", "value1", 10)
	timeMap3.Set("key2", "value2", 20)
	timeMap3.Set("key1", "value1_updated", 30)

	result5 := timeMap3.Get("key1", 15)
	result6 := timeMap3.Get("key1", 35)
	result7 := timeMap3.Get("key2", 25)
	result8 := timeMap3.Get("key2", 5)

	fmt.Printf("Multiple keys test:\n")
	fmt.Printf("  key1 at timestamp 15: \"%s\" (should be \"value1\")\n", result5)
	fmt.Printf("  key1 at timestamp 35: \"%s\" (should be \"value1_updated\")\n", result6)
	fmt.Printf("  key2 at timestamp 25: \"%s\" (should be \"value2\")\n", result7)
	fmt.Printf("  key2 at timestamp 5: \"%s\" (should be empty)\n", result8)
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
