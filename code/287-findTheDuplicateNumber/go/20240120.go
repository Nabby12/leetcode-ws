package main

func findTheDuplicateNumber(nums []int) int {
	numMap := make(map[int]int, len(nums))

	var ans int
	for _, num := range nums {
		if numMap[num] > 0 {
			ans = num
			break
		}
		numMap[num]++
	}

	return ans
}
