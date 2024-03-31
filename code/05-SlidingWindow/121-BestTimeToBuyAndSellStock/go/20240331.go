package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
func bestTimeToBuyAndSellStock(prices []int) int {
	l, r := 0, 1
	maxProfit := 0
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
