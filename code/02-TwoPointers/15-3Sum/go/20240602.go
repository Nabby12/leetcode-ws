package main

import (
	"sort"
)

// https://leetcode.com/problems/3sum/description/
func threeSum8(nums []int) [][]int {
	ans := [][]int{}
	sort.Slice(nums, func(i int, j int) bool { return nums[i] < nums[j] })
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			base, left, right := nums[i], nums[l], nums[r]
			sum := base + left + right
			if sum == 0 {
				ans = append(ans, []int{base, left, right})
				l++
				r--
				for l+1 < len(nums) && nums[l] == nums[l-1] {
					l++
				}
				for r-1 > l && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}
