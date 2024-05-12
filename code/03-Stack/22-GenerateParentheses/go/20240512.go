package main

// https://leetcode.com/problems/generate-parentheses/description/
func generateParentheses9(n int) []string {
	ans := []string{}
	stack := make([]rune, n*2)
	generate10(&ans, stack, n, 0, 0)
	return ans
}

func generate10(ans *[]string, stack []rune, n int, open int, close int) {
	if open+close == n*2 {
		*ans = append(*ans, string(stack))
	}
	if open < n {
		stack[open+close] = '('
		generate10(ans, stack, n, open+1, close)
	}
	if close < open {
		stack[open+close] = ')'
		generate10(ans, stack, n, open, close+1)
	}
}
