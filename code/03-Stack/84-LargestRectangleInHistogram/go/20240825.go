package main

// https://leetcode.com/problems/largest-rectangle-in-histogram/description/
func largestRectangleInHistogram12(heights []int) int {
	maxArea := 0
	stack := [][2]int{}
	for i, height := range heights {
		startIndex := i
		for len(stack) > 0 && height < stack[len(stack)-1][1] {
			lastIndex, lastHeight := stack[len(stack)-1][0], stack[len(stack)-1][1]
			area := lastHeight * (i - lastIndex)
			maxArea = max(maxArea, area)
			startIndex = lastIndex
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, [2]int{startIndex, height})
	}
	for _, el := range stack {
		area := el[1] * (len(heights) - el[0])
		maxArea = max(maxArea, area)
	}
	return maxArea
}
