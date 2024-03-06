package main

import "math"

// https://leetcode.com/problems/koko-eating-bananas/description/
func kokoEatingBananas(piles []int, h int) int {
	var maxPile int
	for _, pile := range piles {
		maxPile = max(maxPile, pile)
	}

	l, r := 1, maxPile
	ans := maxPile
	for l <= r {
		k := (l + r) / 2
		hours := 0
		for _, pile := range piles {
			hours += int(math.Ceil(float64(pile) / float64(k)))
		}

		if hours > h {
			l = k + 1
		} else {
			r = k - 1
			ans = k
		}
	}

	return ans
}
