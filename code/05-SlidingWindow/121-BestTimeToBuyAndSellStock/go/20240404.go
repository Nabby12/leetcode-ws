package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
func bestTimeToBuyAndSellStock3(prices []int) int {
	maxProfit := 0
	l, r := 0, 1
	for r < len(prices) {
		buy, sell := prices[l], prices[r]
		if buy < sell {
			profit := sell - buy
			maxProfit = max(maxProfit, profit)
		} else {
			l = r
		}
		r++
	}
	return maxProfit
}
