package main

// https://leetcode.com/problems/daily-temperatures/description/
func dailyTemperatures7(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := []int{}
	for i, t := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < t {
			lastIndex := stack[len(stack)-1]
			ans[lastIndex] = i - lastIndex
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
