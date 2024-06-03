package main

// https://leetcode.com/problems/trapping-rain-water/description/
func trappingRainWater6(height []int) int {
	ans := 0
	l, r := 0, len(height)-1
	left, right := height[l], height[r]
	for l < r {
		if left <= right {
			l++
			if left < height[l] {
				left = height[l]
				continue
			}
			ans += left - height[l]
		} else {
			r--
			if right < height[r] {
				right = height[r]
				continue
			}
			ans += right - height[r]
		}
	}
	return ans
}
