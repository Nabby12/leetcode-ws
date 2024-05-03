package main

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
func findMinimumInRotatedSortedArray8(nums []int) int {
	l, r := 0, len(nums)-1
	ans := nums[l]
	for l <= r {
		m := (l + r) / 2
		left, mid, right := nums[l], nums[m], nums[r]
		if left < right {
			return min(ans, left)
		}
		ans = min(ans, mid)
		if mid >= left {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return ans
}
