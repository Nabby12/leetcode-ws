package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func longestSubstringWithoutRepeatingCharacters6(s string) int {
	ans := 0
	stack := []rune{}
	strs := []rune(s)
	for _, r := range strs {
		for len(stack) > 0 && isDuplicate6(stack, r) {
			stack = stack[1:]
		}
		stack = append(stack, r)
		ans = max(ans, len(stack))
	}
	return ans
}

func isDuplicate6(runes []rune, r rune) bool {
	for _, v := range runes {
		if v == r {
			return true
		}
	}
	return false
}
