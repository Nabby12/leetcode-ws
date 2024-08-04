package main

// https://leetcode.com/problems/trapping-rain-water/description/
func trappingRainWater7(height []int) int {
	l, r := 0, len(height)-1
	left, right := height[l], height[r]
	trapped := 0
	for l < r {
		if left <= right {
			l++
			if left < height[l] {
				left = height[l]
				continue
			}
			trapped += left - height[l]
		} else {
			r--
			if right < height[r] {
				right = height[r]
				continue
			}
			trapped += right - height[r]
		}
	}
	return trapped
}
