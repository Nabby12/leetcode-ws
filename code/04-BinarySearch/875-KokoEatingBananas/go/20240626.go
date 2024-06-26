package main

import (
	"sort"
)

// https://leetcode.com/problems/koko-eating-bananas/description/
func kokoEatingBananas5(piles []int, h int) int {
	sort.Slice(piles, func(i, j int) bool {
		return piles[i] > piles[j]
	})
	total := 0
	for _, pile := range piles {
		total += pile
	}
	minK := 1
	maxK := piles[0]
	k := maxK
	for minK <= maxK {
		midK := (minK + maxK) / 2
		hours := 0
		for _, pile := range piles {
			hour := pile / midK
			if pile%midK > 0 {
				hour++
			}
			hours += hour
		}
		if hours <= h {
			maxK = midK - 1
			k = midK
		} else {
			minK = midK + 1
		}
	}
	return k
}
