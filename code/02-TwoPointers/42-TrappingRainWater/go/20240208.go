package main

// https://leetcode.com/problems/trapping-rain-water/description/
func trappingRainWater(height []int) int {
	trapAmount := 0
	l, r := 0, len(height)-1
	maxLeftHeight, maxRightHeight := height[l], height[r]

	for l < r {
		if maxLeftHeight < maxRightHeight {
			l++
			if maxLeftHeight < height[l] {
				maxLeftHeight = height[l]
				continue
			}
			trapAmount += maxLeftHeight - height[l]
		} else {
			r--
			if maxRightHeight < height[r] {
				maxRightHeight = height[r]
				continue
			}
			trapAmount += maxRightHeight - height[r]
		}
	}

	return trapAmount
}
