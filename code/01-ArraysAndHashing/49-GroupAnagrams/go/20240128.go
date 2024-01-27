package main

import "sort"

// https://leetcode.com/problems/group-anagrams/description/
func groupAnagram2(strs []string) [][]string {
	hashMap := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortStr(str)
		hashMap[sortedStr] = append(hashMap[sortedStr], str)
	}

	ans := [][]string{}
	for _, values := range hashMap {
		ans = append(ans, values)
	}

	return ans
}

func sortStr(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
