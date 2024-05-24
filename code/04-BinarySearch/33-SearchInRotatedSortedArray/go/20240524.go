package main

// https://leetcode.com/problems/search-in-rotated-sorted-array/description/
func searchInRotatedSortedArray5(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		left, mid, right := nums[l], nums[m], nums[r]
		if target == mid {
			return m
		}
		if left <= mid {
			if target < left || mid < target {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if target < mid || right < target {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return -1
}
