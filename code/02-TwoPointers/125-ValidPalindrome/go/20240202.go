package main

import (
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome/description/
func validPalindrome(s string) bool {
	r := []rune(s)
	var converted string
	for _, v := range r {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			lower := unicode.ToLower(v)
			converted += string(lower)
		}
	}

	for i, _ := range converted {
		j := len(converted) - 1 - i
		if converted[i] != converted[j] {
			return false
		}
	}

	return true
}
