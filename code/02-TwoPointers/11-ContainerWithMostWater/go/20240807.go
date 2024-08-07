package main

// https://leetcode.com/problems/container-with-most-water/description/
func containerWithMostWater4(height []int) int {
	maxArea := 0
	l, r := 0, len(height)-1
	for l < r {
		area := 0
		if height[l] < height[r] {
			area = height[l] * (r - l)
			l++
		} else {
			area = height[r] * (r - l)
			r--
		}
		maxArea = max(maxArea, area)
	}
	return maxArea
}
