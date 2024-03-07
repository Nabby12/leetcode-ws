package main

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/koko-eating-bananas/description/
func kokoEatingBananas2(piles []int, h int) int {
	sort.Ints(piles)
	maxPile := piles[len(piles)-1]

	minK, maxK := 1, maxPile
	ans := maxPile
	for minK <= maxK {
		k := (minK + maxK) / 2
		hours := 0
		for _, pile := range piles {
			hours += int(math.Ceil(float64(pile) / float64(k)))
		}

		if hours > h {
			minK = k + 1
		} else {
			maxK = k - 1
			ans = k
		}
	}

	return ans
}
