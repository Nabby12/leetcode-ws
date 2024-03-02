package main

// https://leetcode.com/problems/binary-search/description/
func binarySearch2(nums []int, target int) int {
	length := len(nums)
	l, r := 0, length-1

	for l <= r {
		mid := (l + r) / 2
		midVal := nums[mid]
		if midVal == target {
			return mid
		} else if midVal < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
