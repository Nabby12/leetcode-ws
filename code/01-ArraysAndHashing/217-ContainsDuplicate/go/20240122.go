package main

// https://leetcode.com/problems/contains-duplicate/description/
func containsDuplicate(nums []int) bool {
	ans := false
	numMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		if numMap[num] {
			ans = true
			break
		}
		numMap[num] = true
	}
	return ans
}
