package main

// https://leetcode.com/problems/valid-anagram/description/
func validAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ans := true
	sMap := make(map[rune]int, len(s))
	for _, v := range s {
		sMap[v]++
	}
	for _, v := range t {
		count, ok := sMap[v]
		if count == 0 || !ok {
			ans = false
			break
		}
		sMap[v]--
	}
	return ans
}
