package main

// https://leetcode.com/problems/generate-parentheses/description/
func generateParentheses3(n int) []string {
	ans := []string{}
	currentStack := make([]rune, n*2)
	generate(&ans, n, 0, 0, currentStack)
	return ans
}

func generate(ansPtr *[]string, n int, open int, close int, currentStack []rune) {
	if len(currentStack) > 0 && n == open && n == close {
		*ansPtr = append(*ansPtr, string(currentStack))
		return
	}
	if open < n {
		currentStack[open+close] = '('
		generate(ansPtr, n, open+1, close, currentStack)
	}
	if close < open {
		currentStack[open+close] = ')'
		generate(ansPtr, n, open, close+1, currentStack)
	}
}
