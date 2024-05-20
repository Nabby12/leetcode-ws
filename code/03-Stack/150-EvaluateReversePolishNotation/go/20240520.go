package main

import "strconv"

// https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
func evaluateReversePolishNotation5(tokens []string) int {
	evaluater := map[string]func(int, int) int{
		"+": func(num1 int, num2 int) int { return num1 + num2 },
		"-": func(num1 int, num2 int) int { return num1 - num2 },
		"*": func(num1 int, num2 int) int { return num1 * num2 },
		"/": func(num1 int, num2 int) int { return num1 / num2 },
	}
	ans := []int{}
	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			ans = append(ans, num)
		} else {
			calc := evaluater[token](ans[len(ans)-2], ans[len(ans)-1])
			ans = ans[:len(ans)-1]
			ans[len(ans)-1] = calc
		}
	}
	return ans[0]
}
