package main

// https://leetcode.com/problems/trapping-rain-water/description/
func trappingRainWater5(height []int) int {
	amount := 0
	l, r := 0, len(height)-1
	left, right := height[l], height[r]
	for l < r {
		if left < right {
			l++
			if left < height[l] {
				left = height[l]
				continue
			}
			amount += left - height[l]
		} else {
			r--
			if right < height[r] {
				right = height[r]
				continue
			}
			amount += right - height[r]
		}
	}
	return amount
}
