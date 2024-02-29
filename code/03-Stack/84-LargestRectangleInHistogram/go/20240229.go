package main

// https://leetcode.com/problems/largest-rectangle-in-histogram/description/
func largestRectangleInHistogram(heights []int) int {
	maxArea := 0
	stack := [][2]int{}

	for i, height := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1][1] > height {
			lastElement := stack[len(stack)-1]
			area := lastElement[1] * (i - lastElement[0])
			maxArea = max(maxArea, area)
			start = lastElement[0]
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, [2]int{start, height})
	}

	for _, v := range stack {
		area := v[1] * (len(heights) - v[0])
		maxArea = max(maxArea, area)
	}

	return maxArea
}
