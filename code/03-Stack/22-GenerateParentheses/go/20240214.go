package main

// https://leetcode.com/problems/generate-parentheses/description/
func generateParentheses2(n int) []string {
	ans := []string{}
	currentStack := make([]rune, n*2)
	generate(&ans, n, 0, 0, currentStack)
	return ans
}

func generate(strs *[]string, n int, open int, close int, currentStack []rune) {
	if open == n && close == n {
		*strs = append(*strs, string(currentStack))
		return
	}
	if open < n {
		currentStack[open+close] = '('
		generate(strs, n, open+1, close, currentStack)
	}
	if close < open {
		currentStack[open+close] = ')'
		generate(strs, n, open, close+1, currentStack)
	}
}
