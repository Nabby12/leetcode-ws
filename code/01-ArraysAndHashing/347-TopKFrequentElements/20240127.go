package main

// https://leetcode.com/problems/top-k-frequent-elements/description/
func topKFrequentElements(nums []int, k int) []int {
	numMap := make(map[int]int, len(nums))
	for _, num := range nums {
		numMap[num]++
	}

	ans := []int{}

	freq := make(map[int][]int)
	for key, num := range numMap {
		freq[num] = append(freq[num], key)
	}

	for i := len(nums); i > 0; i-- {
		if num, ok := freq[i]; ok {
			ans = append(ans, num...)
			if len(ans) >= k {
				ans = ans[:k]
				break
			}
		}
	}

	return ans
}
