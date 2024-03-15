package main

// https://leetcode.com/problems/search-a-2d-matrix/description/
func searchA2DMatrix6(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])

	targetRowIndex := -1
	t, b := 0, rows-1
	for t <= b {
		row := (t + b) / 2
		first, last := matrix[row][0], matrix[row][cols-1]
		if target > last {
			t = row + 1
		} else if target < first {
			b = row - 1
		} else {
			targetRowIndex = row
			break
		}
	}

	if targetRowIndex == -1 {
		return false
	}

	targetRow := matrix[targetRowIndex]
	l, r := 0, cols-1
	for l <= r {
		m := (l + r) / 2
		mid := targetRow[m]
		if target > mid {
			l = m + 1
		} else if target < mid {
			r = m - 1
		} else {
			return true
		}
	}

	return false
}
