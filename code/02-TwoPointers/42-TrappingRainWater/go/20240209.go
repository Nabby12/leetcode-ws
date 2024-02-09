package main

// https://leetcode.com/problems/trapping-rain-water/description/
func trappingRainWater2(height []int) int {
	l, r := 0, len(height)-1
	left, right := height[l], height[r]

	trapAmount := 0
	for l < r {
		if left < right {
			l++
			if left < height[l] {
				left = height[l]
				continue
			}
			trapAmount += left - height[l]
		} else {
			r--
			if right < height[r] {
				right = height[r]
				continue
			}
			trapAmount += right - height[r]
		}
	}

	return trapAmount
}
