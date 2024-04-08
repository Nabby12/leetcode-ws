package main

// https://leetcode.com/problems/daily-temperatures/description/
func dailyTemperatures3(temperatures []int) []int {
	stack := []int{}
	ans := make([]int, len(temperatures))
	for i, temperature := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperature {
			index := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			ans[index] = i - index
		}
		stack = append(stack, i)
	}
	return ans
}
