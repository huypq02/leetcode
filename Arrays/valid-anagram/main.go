package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// s = ls[b,a,c]
	// t = lt[c,b,a]
	var ls = make(map[rune]int)
	for _, v := range s {
		ls[v]++
	}
	var lt = make(map[rune]int)
	for _, v := range t {
		lt[v]++
	}

	// comparing two lists
	for k, vs := range ls {
		if vk, ok := lt[k]; !ok || vs != vk {
			return false
		}
	}

	return true
}

func main() {
	s1, t1 := "anagram", "nagaram"
	println(isAnagram(s1, t1)) // output: true

	s2, t2 := "rat", "car"
	println(isAnagram(s2, t2)) // output: false

	s3, t3 := "chán", "nán"
	println(isAnagram(s3, t3)) // output: false

	s4, t4 := "ngân", "nâng"
	println(isAnagram(s4, t4)) // output: true
}
