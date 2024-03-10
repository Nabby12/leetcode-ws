package main

// https://leetcode.com/problems/search-in-rotated-sorted-array/description/
func searchInRotatedSortedArray(nums []int, target int) int {
	ans := -1
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		mid := nums[m]
		if mid == target {
			ans = m
			break
		}
		if nums[l] <= mid {
			if target > mid || target < nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if target < mid || target > nums[r] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return ans
}
