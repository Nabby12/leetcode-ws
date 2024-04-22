package main

// https://leetcode.com/problems/generate-parentheses/description/
func generateParentheses6(n int) []string {
	ans := []string{}
	stack := make([]rune, n*2)
	generate6(&ans, n, 0, 0, stack)
	return ans
}

func generate6(ans *[]string, n int, open int, close int, stack []rune) {
	if len(stack) > 0 && open == n && close == n {
		*ans = append(*ans, string(stack))
	}
	if open < n {
		stack[open+close] = '('
		generate6(ans, n, open+1, close, stack)
	}
	if close < open {
		stack[open+close] = ')'
		generate6(ans, n, open, close+1, stack)
	}
}
