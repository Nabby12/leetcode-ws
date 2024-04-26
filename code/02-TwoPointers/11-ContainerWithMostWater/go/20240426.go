package main

// https://leetcode.com/problems/container-with-most-water/description/
func containerWithMostWater2(height []int) int {
	ans := 0
	l, r := 0, len(height)-1
	for l < r {
		left, right := height[l], height[r]
		amount := 0
		width := r - l
		if left <= right {
			amount = left * width
			l++
		} else {
			amount = right * width
			r--
		}
		ans = max(ans, amount)
	}
	return ans
}
