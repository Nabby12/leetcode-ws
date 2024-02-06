package main

// https://leetcode.com/problems/product-of-array-except-self/description/
func productOfArrayExceptSelf2(nums []int) []int {
	length := len(nums)
	left := make(map[int]int, length)
	left[0] = 1
	right := make(map[int]int, length)
	right[length-1] = 1
	ans := []int{}

	for i := 1; i < length; i++ {
		l := i
		r := length - 1 - i
		left[l] = nums[l-1] * left[l-1]
		right[r] = nums[r+1] * right[r+1]
	}

	for i, _ := range nums {
		ans = append(ans, left[i]*right[i])
	}

	return ans
}
