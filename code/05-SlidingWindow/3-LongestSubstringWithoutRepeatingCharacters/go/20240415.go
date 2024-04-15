package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func longestSubstringWithoutRepeatingCharacters3(s string) int {
	l, r := 0, 0
	strs := []rune(s)
	buf := []rune{}
	ans := 0
	for _, v := range strs {
		for len(buf) > 0 && contains3(buf, v) {
			buf = buf[1:]
			l++
		}
		buf = append(buf, v)
		ans = max(ans, len(buf))
		r++
	}
	return ans
}

func contains3(runes []rune, target rune) bool {
	for _, r := range runes {
		if r == target {
			return true
		}
	}
	return false
}
