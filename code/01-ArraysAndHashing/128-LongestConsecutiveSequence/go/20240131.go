package main

// https://leetcode.com/problems/longest-consecutive-sequence/
func longestConsecutiveSequence(nums []int) int {
	numMap := map[int]bool{}
	for _, num := range nums {
		numMap[num] = true
	}

	ans := 0
	for num, _ := range numMap {
		if numMap[num-1] {
			continue
		}

		count, nextNum := 1, num+1
		for numMap[nextNum] {
			count++
			nextNum++
		}

		ans = max(ans, count)

		if count > len(nums)/2 {
			break
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
