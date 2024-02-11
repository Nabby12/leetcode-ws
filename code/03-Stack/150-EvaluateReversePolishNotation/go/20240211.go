package main

import (
	"strconv"
)

// https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
func evaluateReversePolishNotation(tokens []string) int {
	numStack := []int{}
	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err != nil {
			eval := evaluate(token, numStack[len(numStack)-2], numStack[len(numStack)-1])
			numStack = numStack[0 : len(numStack)-2]
			numStack = append(numStack, eval)
		} else {
			numStack = append(numStack, num)
		}
	}

	return numStack[0]
}

func evaluate(operand string, num1 int, num2 int) int {
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
