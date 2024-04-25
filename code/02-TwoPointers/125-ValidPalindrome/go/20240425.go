package main

import (
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome/description/
func validPalindrome4(s string) bool {
	runes := []rune{}
	for _, r := range []rune(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			runes = append(runes, unicode.ToLower(r))
		}
	}
	l, r := 0, len(runes)-1
	for l < r {
		left, right := runes[l], runes[r]
		if left != right {
			return false
		}
		l++
		r--
	}
	return true
}
