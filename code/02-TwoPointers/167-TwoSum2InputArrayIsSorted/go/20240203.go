package main

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
func twoSum2InputArrayIsSorted(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l <= r {
		sum := numbers[l] + numbers[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return []int{}
}
