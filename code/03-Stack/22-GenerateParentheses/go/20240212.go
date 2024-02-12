package main

// https://leetcode.com/problems/generate-parentheses/description/
func generateParentheses(n int) []string {
	ans := []string{}
	stack := make([]rune, n*2)
	generate(&ans, n, 0, 0, stack)
	return ans
}

func generate(ans *[]string, n int, open int, close int, stack []rune) {
	if open == n && close == n {
		*ans = append(*ans, string(stack))
		return
	}
	if open < n {
		stack[open+close] = '('
		generate(ans, n, open+1, close, stack)
	}
	if close < open {
		stack[open+close] = ')'
		generate(ans, n, open, close+1, stack)
	}
}
