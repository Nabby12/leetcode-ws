package main

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
func findMinimumInRotatedSortedArray2(nums []int) int {
	ans := nums[0]
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] < nums[r] {
			ans = min(ans, nums[l])
			break
		}
		m := (l + r) / 2
		mid := nums[m]
		ans = min(ans, mid)
		if nums[l] <= mid {
			l = m + 1
		} else if nums[r] > mid {
			r = m - 1
		}
	}
	return ans
}
