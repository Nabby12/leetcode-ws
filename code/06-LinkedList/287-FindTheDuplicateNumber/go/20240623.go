package main

// https://leetcode.com/problems/find-the-duplicate-number/description/
func findTheDuplicateNumber2(nums []int) int {
	numMap := map[int]bool{}
	for _, num := range nums {
		if numMap[num] {
			return num
		}
		numMap[num] = true
	}
	return -1
}
