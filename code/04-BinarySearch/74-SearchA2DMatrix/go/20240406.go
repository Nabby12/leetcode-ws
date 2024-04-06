package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
func searchA2DMatrix7(matrix [][]int, target int) bool {
	t, b := 0, len(matrix)-1
	targetRow := []int{}
	for t <= b {
		m := (t + b) / 2
		row := matrix[m]
		first, last := row[0], row[len(row)-1]
		if first <= target && target <= last {
			targetRow = row
			break
		}
		if target > last {
			t = m + 1
		} else {
			b = m - 1
		}
	}
	if len(targetRow) == 0 {
		return false
	}
	l, r := 0, len(targetRow)-1
	for l <= r {
		m := (l + r) / 2
		mid := targetRow[m]
		if target == mid {
			return true
		}
		if target > mid {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}
