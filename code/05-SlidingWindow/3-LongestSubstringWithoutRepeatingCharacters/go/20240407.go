package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func longestSubstringWithoutRepeatingCharacters(s string) int {
	strs := []rune(s)
	buf := []rune{}
	ans := 0
	l, r := 0, 0
	for r < len(strs) {
		right := strs[r]
		for len(buf) > 0 && contains(buf, right) {
			buf = buf[1:]
			l++
		}
		buf = append(buf, right)
		ans = max(ans, len(buf))
		r++
	}
	return ans
}

func contains(runes []rune, target rune) bool {
	for _, s := range runes {
		if s == target {
			return true
		}
	}
	return false
}
