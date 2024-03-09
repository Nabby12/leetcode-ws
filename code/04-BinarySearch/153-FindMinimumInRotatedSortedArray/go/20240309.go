package main

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
func findMinimumInRotatedSortedArray(nums []int) int {
	ans := nums[0]
	l, r := 0, len(nums)-1

	for l <= r {
		if nums[l] < nums[r] {
			ans = min(ans, nums[l])
			break
		}

		m := (l + r) / 2
		ans = min(ans, nums[m])
		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return ans
}
