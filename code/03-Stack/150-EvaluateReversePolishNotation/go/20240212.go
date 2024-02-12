package main

import "strconv"

// https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
func evaluateReversePolishNotation2(tokens []string) int {
	numStack := []int{}
	evalFuncs := map[string]func(num1 int, num2 int) int{
		"+": func(num1 int, num2 int) int { return num1 + num2 },
		"-": func(num1 int, num2 int) int { return num1 - num2 },
		"*": func(num1 int, num2 int) int { return num1 * num2 },
		"/": func(num1 int, num2 int) int { return num1 / num2 },
	}

	for _, token := range tokens {
		if evalFunc, ok := evalFuncs[token]; ok {
			eval := evalFunc(numStack[len(numStack)-2], numStack[len(numStack)-1])
			numStack = numStack[0 : len(numStack)-2]
			numStack = append(numStack, eval)
		} else {
			if num, err := strconv.Atoi(token); err == nil {
				numStack = append(numStack, num)
			}
		}
	}

	return numStack[0]
}
