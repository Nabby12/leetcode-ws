package main

// https://leetcode.com/problems/valid-parentheses/description/
func validParentheses2(s string) bool {
	if len(s) < 2 {
		return false
	}

	closure := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}
	for _, r := range []rune(s) {
		if _, ok := closure[r]; ok {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return false
			}
			if r == closure[stack[len(stack)-1]] {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
