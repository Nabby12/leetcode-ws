package main

import (
	"strings"
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome/description/
func validPalindrome3(s string) bool {
	converter := func(s rune) rune {
		if !unicode.IsLetter(s) && !unicode.IsDigit(s) {
			return -1
		}
		return unicode.ToLower(s)
	}
	converted := strings.Map(converter, s)

	index, reverseIndex := 0, len(converted)-1
	for index < reverseIndex {
		if converted[index] != converted[reverseIndex] {
			return false
		}
		index++
		reverseIndex--
	}
	return true
}
