package main

// https://leetcode.com/problems/product-of-array-except-self/description/
func productOfArrayExceptSelf(nums []int) []int {
	length := len(nums)
	left := make([]int, length)
	right := make([]int, length)
	ans := make([]int, length)
	left[0] = 1
	right[length-1] = 1

	for i := 1; i < length; i++ {
		l := i
		r := length - i - 1
		left[l] = nums[l-1] * left[l-1]
		right[r] = nums[r+1] * right[r+1]
	}

	for i := 0; i < length; i++ {
		ans[i] = left[i] * right[i]
	}

	return ans
}
