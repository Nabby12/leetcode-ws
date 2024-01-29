package main

import "fmt"

// https://leetcode.com/problems/top-k-frequent-elements/description/
func topKFrequentElements2(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	freq := make(map[int][]int)
	for num, count := range countMap {
		freq[count] = append(freq[count], num)
	}

	ans := []int{}
	for i := len(nums); i > 0; i-- {
		if values, ok := freq[i]; ok {
			fmt.Println(values)
			ans = append(ans, values...)
		}
		if len(ans) >= k {
			break
		}
	}

	return ans
}
