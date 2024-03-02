package main

// https://leetcode.com/problems/largest-rectangle-in-histogram/description/
func largestRectangleInHistogram3(heights []int) int {
	maxArea := 0
	stack := [][2]int{}

	for i, h := range heights {
		start := i

		for len(stack) > 0 && h < stack[len(stack)-1][1] {
			lastIndex, lastHeight := stack[len(stack)-1][0], stack[len(stack)-1][1]
			area := lastHeight * (i - lastIndex)
			maxArea = max(maxArea, area)
			stack = stack[0 : len(stack)-1]
		}

		stack = append(stack, [2]int{start, h})
	}

	for _, el := range stack {
		area := el[1] * (len(heights) - el[0])
		maxArea = max(maxArea, area)
		stack = stack[0 : len(stack)-1]
	}

	return maxArea
}
