package main

// https://leetcode.com/problems/koko-eating-bananas/description/
func kokoEatingBananas4(piles []int, h int) int {
	var maxPile int
	for _, pile := range piles {
		maxPile = max(maxPile, pile)
	}
	l, r := 1, maxPile
	ans := r
	for l <= r {
		m := (l + r) / 2
		hours := 0
		for _, pile := range piles {
			hour := pile / m
			if pile%m > 0 {
				hour++
			}
			hours += hour
		}
		if hours > h {
			l = m + 1
		} else {
			r = m - 1
			ans = m
		}
	}
	return ans
}
