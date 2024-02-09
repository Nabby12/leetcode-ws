package main

// https://leetcode.com/problems/valid-parentheses/description/
func validParentheses(s string) bool {
	if len(s) < 2 {
		return false
	}

	stack := []rune{}
	closure := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, r := range s {
		if _, ok := closure[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || closure[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[0 : len(stack)-1]
		}
	}

	return len(stack) == 0
}
