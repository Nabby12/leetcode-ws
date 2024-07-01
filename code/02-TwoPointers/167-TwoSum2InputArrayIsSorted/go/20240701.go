package main

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
func twoSum2InputArrayIsSorted2(numbers []int, target int) []int {
	ans := make([]int, 2)
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			ans = []int{l + 1, r + 1}
			break
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return ans
}
