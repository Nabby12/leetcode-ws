package main

// https://leetcode.com/problems/generate-parentheses/description/
func generateParentheses7(n int) []string {
	ans := []string{}
	stack := make([]rune, n*2)
	generate7(&ans, stack, n, 0, 0)
	return ans
}

func generate7(ans *[]string, stack []rune, n int, open int, close int) {
	if open == n && close == n {
		*ans = append(*ans, string(stack))
	}
	if open < n {
		stack[open+close] = '('
		generate7(ans, stack, n, open+1, close)
	}
	if close < open {
		stack[open+close] = ')'
		generate7(ans, stack, n, open, close+1)
	}
}
