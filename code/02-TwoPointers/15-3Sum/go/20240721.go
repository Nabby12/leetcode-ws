package main

import "sort"

// https://leetcode.com/problems/3sum/description/
func threeSum10(nums []int) [][]int {
	ans := [][]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i, num := range nums {
		if i > len(nums)-3 {
			break
		}
		if i > 0 && num == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := num + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{num, nums[l], nums[r]})
				l++
				r--
				for l < len(nums) && nums[l] == nums[l-1] {
					l++
				}
				for r > l && nums[r] == nums[r+1] {
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
