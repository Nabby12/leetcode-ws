package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func longestSubstringWithoutRepeatingCharacters4(s string) int {
	strs := []rune(s)
	stack := []rune{}
	ans := 0
	for _, str := range strs {
		for len(stack) > 0 && contains4(stack, str) {
			stack = stack[1:]
		}
		stack = append(stack, str)
		ans = max(ans, len(stack))
	}
	return ans
}

func contains4(runes []rune, target rune) bool {
	for _, r := range runes {
		if r == target {
			return true
		}
	}
	return false
}
