package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
func bestTimeToBuyAndSellStock5(prices []int) int {
	ans := 0
	l, r := 0, 1
	for l < r && r < len(prices) {
		if prices[l] < prices[r] {
			ans = max(ans, prices[r]-prices[l])
		} else {
			l = r
		}
		r++
	}
	return ans
}
