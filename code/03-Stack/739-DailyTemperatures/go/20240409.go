package main

// https://leetcode.com/problems/daily-temperatures/description/
func dailyTemperatures4(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := []int{}
	for i, temperature := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperature {
			lastIndex := stack[len(stack)-1]
			days := i - lastIndex
			ans[lastIndex] = days
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
