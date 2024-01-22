package main

import (
	"sort"
)

// https://leetcode.com/problems/group-anagrams/description/
func groupAnagram(strs []string) [][]string {
	hashMap := make(map[string][]string)
	for _, str := range strs {
		sorted := sortString(str)
		hashMap[sorted] = append(hashMap[sorted], str)
	}

	var ans [][]string
	for _, values := range hashMap {
		ans = append(ans, values)
	}
	return ans
}

func sortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
