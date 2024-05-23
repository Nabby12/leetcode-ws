package main

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
func findMinimumInRotatedSortedArray9(nums []int) int {
	ans := nums[0]
	l, r := 0, len(nums)-1
	for l <= r {
		left, right := nums[l], nums[r]
		if left < right {
			ans = min(ans, left)
			break
		}
		m := (l + r) / 2
		mid := nums[m]
		ans = min(ans, mid)
		if left <= mid {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return ans
}
