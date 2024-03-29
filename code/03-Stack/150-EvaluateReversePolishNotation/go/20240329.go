package main

import "strconv"

// https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
func evaluateReversePolishNotation3(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
			continue
		}
		res := evaluate3(stack[len(stack)-2], stack[len(stack)-1], token)
		stack = stack[0 : len(stack)-2]
		stack = append(stack, res)
	}
	return stack[len(stack)-1]
}

func evaluate3(num1 int, num2 int, operand string) int {
	var ans int
	switch operand {
	case "+":
		ans = num1 + num2
	case "-":
		ans = num1 - num2
	case "*":
		ans = num1 * num2

	case "/":
		ans = num1 / num2
	}
	return ans
}
