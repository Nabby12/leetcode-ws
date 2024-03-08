package main

import "math"

// https://leetcode.com/problems/koko-eating-bananas/description/
func kokoEatingBananas3(piles []int, h int) int {
	var maxPile int
	for _, pile := range piles {
		maxPile = max(maxPile, pile)
	}

	var ans int
	min, max := 1, maxPile
	for min <= max {
		k := (min + max) / 2
		hours := 0
		for _, pile := range piles {
			hours += int(math.Ceil(float64(pile) / float64(k)))
		}

		if hours > h {
			min = k + 1
		} else {
			max = k - 1
			ans = k
		}
	}

	return ans
}
