package main

// https://leetcode.com/problems/search-a-2d-matrix/description/
func searchA2DMatrix8(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	t, b := 0, rows-1
	targetIndex := -1
	for t <= b {
		m := (t + b) / 2
		row := matrix[m]
		left, right := row[0], row[cols-1]
		if target < left {
			b = m - 1
		} else if target > right {
			t = m + 1
		} else {
			targetIndex = m
			break
		}
	}
	if targetIndex == -1 {
		return false
	}
	l, r := 0, cols-1
	targetRow := matrix[targetIndex]
	for l <= r {
		m := (l + r) / 2
		mid := targetRow[m]
		if target < mid {
			r = m - 1
		} else if target > mid {
			l = m + 1
		} else {
			return true
		}
	}
	return false
}
