package main

// https://leetcode.com/problems/valid-anagram/description/
func validAnagram20240127(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := 0
	sMap := make(map[rune]int, len(s))
	for _, str := range s {
		sMap[str]++
		count++
	}
	for _, str := range t {
		if sMap[str] > 0 {
			sMap[str]--
			count--
		}
	}

	if count == 0 {
		return true
	}

	return false
}
