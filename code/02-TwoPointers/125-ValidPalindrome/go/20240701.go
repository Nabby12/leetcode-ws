package main

import (
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome/description/
func validPalindrome5(s string) bool {
	stack := []rune{}
	for _, r := range []rune(s) {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			stack = append(stack, unicode.ToLower(r))
		}
	}
	l, r := 0, len(stack)-1
	for l <= r {
		left, right := stack[l], stack[r]
		if left != right {
			return false
		}
		l++
		r--
	}
	return true
}
