package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func longestSubstringWithoutRepeatingCharacters5(s string) int {
	stack := []rune{}
	strs := []rune(s)
	maxLength := 0
	for _, str := range strs {
		for len(stack) > 0 && contains5(stack, str) {
			stack = stack[1:len(stack)]
		}
		stack = append(stack, str)
		maxLength = max(maxLength, len(stack))
	}
	return maxLength
}

func contains5(runes []rune, target rune) bool {
	for _, r := range runes {
		if r == target {
			return true
		}
	}
	return false
}
