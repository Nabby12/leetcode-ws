package main

// https://leetcode.com/problems/search-in-rotated-sorted-array/description/
func searchInRotatedSortedArray4(nums []int, target int) int {
	ans := -1
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		mid := nums[m]
		if mid == target {
			return m
		}
		left, right := nums[l], nums[r]
		if left <= mid {
			if target > mid || target < left {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if target < mid || target > right {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return ans
}
