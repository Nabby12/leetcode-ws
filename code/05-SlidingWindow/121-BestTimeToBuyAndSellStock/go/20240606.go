package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
func bestTimeToBuyAndSellStock4(prices []int) int {
	ans := 0
	l := 0
	for i := 0; i < len(prices); i++ {
		if prices[l] > prices[i] {
			l = i
		}
		ans = max(ans, prices[i]-prices[l])
	}
	return ans
}
