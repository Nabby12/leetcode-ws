package main

import (
	"strings"
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome/description/
func validPalindrome2(s string) bool {
	convert := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return -1
		}
		return unicode.ToLower(r)
	}
	converted := strings.Map(convert, s)

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
