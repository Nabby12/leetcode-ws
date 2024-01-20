package main

// https://leetcode.com/problems/binary-search/description/
func binarySearch(nums []int, target int) int {
	ans := -1

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		midVal := nums[mid]

		if target == midVal {
			ans = mid
			break
		} else if target > midVal {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}
