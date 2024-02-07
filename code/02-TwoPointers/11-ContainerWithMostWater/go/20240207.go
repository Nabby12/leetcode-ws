package main

// https://leetcode.com/problems/container-with-most-water/description/
func containerWithMostWater(height []int) int {
	l, r := 0, len(height)-1
	area := 0

	for l < r {
		width := r - l
		if height[l] < height[r] {
			area = max(area, height[l]*width)
			l++
		} else {
			area = max(area, height[r]*width)
			r--
		}
	}

	return area
}
