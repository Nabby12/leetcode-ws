package main

import "strconv"

// https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
func evaluateReversePolishNotation4(tokens []string) int {
	evaluation := map[string]func(int, int) int{
		"+": func(num1 int, num2 int) int { return num1 + num2 },
		"-": func(num1 int, num2 int) int { return num1 - num2 },
		"*": func(num1 int, num2 int) int { return num1 * num2 },
		"/": func(num1 int, num2 int) int { return num1 / num2 },
	}
	stack := []int{}
	for _, v := range tokens {
		if num, ok := strconv.Atoi(v); ok == nil {
			stack = append(stack, num)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			calc := evaluation[v](num1, num2)
			stack = stack[0 : len(stack)-2]
			stack = append(stack, calc)
		}
	}
	return stack[len(stack)-1]
}
