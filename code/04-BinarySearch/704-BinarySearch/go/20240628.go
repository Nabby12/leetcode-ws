package main

// https://leetcode.com/problems/binary-search/description/
func binarySearch4(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		mid := nums[m]
		if mid < target {
			l = m + 1
		} else if mid > target {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}
