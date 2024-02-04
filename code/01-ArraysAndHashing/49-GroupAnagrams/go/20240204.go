package main

import (
	"sort"
)

// https://leetcode.com/problems/group-anagrams/description/
func groupAnagram3(strs []string) [][]string {
	hashMap := map[string][]string{}
	for _, str := range strs {
		sorted := sortedString(str)
		hashMap[sorted] = append(hashMap[sorted], str)
	}

	ans := [][]string{}
	for _, strs := range hashMap {
		ans = append(ans, strs)
	}

	return ans
}

func sortedString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i int, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
