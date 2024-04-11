package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func longestSubstringWithoutRepeatingCharacters2(s string) int {
	l, r := 0, 0
	strs := []rune(s)
	buf := []rune{}
	maxLen := 0
	for r < len(s) {
		right := strs[r]
		for len(strs) > 0 && contains2(buf, right) {
			buf = buf[1:]
			l++
		}
		buf = append(buf, right)
		maxLen = max(maxLen, len(buf))
		r++
	}
	return maxLen
}

func contains2(runes []rune, target rune) bool {
	for _, r := range runes {
		if r == target {
			return true
		}
	}
	return false
}
