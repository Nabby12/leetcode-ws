package main

// https://leetcode.com/problems/container-with-most-water/description
func containerWithMostWater3(height []int) int {
	maxArea := 0
	l, r := 0, len(height)-1
	for l < r {
		left, right := height[l], height[r]
		width := r - l
		if left > right {
			area := right * width
			maxArea = max(maxArea, area)
			r--
		} else {
			area := left * width
			maxArea = max(maxArea, area)
			l++
		}
	}
	return maxArea
}
