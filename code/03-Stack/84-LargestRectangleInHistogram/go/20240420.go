package main

// https://leetcode.com/problems/largest-rectangle-in-histogram/description/
func largestRectangleInHistogram10(heights []int) int {
	maxArea := 0
	stack := [][2]int{}
	for i, h := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1][1] > h {
			lastIndex, lastHeight := stack[len(stack)-1][0], stack[len(stack)-1][1]
			area := lastHeight * (i - lastIndex)
			maxArea = max(maxArea, area)
			stack = stack[0 : len(stack)-1]
			start = lastIndex
		}
		stack = append(stack, [2]int{start, h})
	}
	for _, v := range stack {
		area := v[1] * (len(heights) - v[0])
		maxArea = max(maxArea, area)
	}
	return maxArea
}
