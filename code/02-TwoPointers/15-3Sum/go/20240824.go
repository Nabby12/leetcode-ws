package main

import (
	"sort"
)

// https://leetcode.com/problems/3sum/description/
func threeSum11(nums []int) [][]int {
	ans := [][]int{}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l-1] == nums[l] {
					l++
				}
				r--
				for r > l && nums[r+1] == nums[r] {
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
