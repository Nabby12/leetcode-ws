package main

// https://leetcode.com/problems/generate-parentheses/description/
func generateParentheses8(n int) []string {
	ans := []string{}
	stack := make([]rune, n*2)
	generate8(&ans, n, 0, 0, stack)
	return ans
}

func generate8(ans *[]string, n int, open int, close int, stack []rune) {
	if open+close == n*2 {
		*ans = append(*ans, string(stack))
		return
	}
	if open < n {
		stack[open+close] = '('
		generate8(ans, n, open+1, close, stack)
	}
	if close < open {
		stack[open+close] = ')'
		generate8(ans, n, open, close+1, stack)
	}
}
