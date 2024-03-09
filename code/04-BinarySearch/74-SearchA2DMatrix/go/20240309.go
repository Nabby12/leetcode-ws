package main

// https://leetcode.com/problems/search-a-2d-matrix/description/
func searchA2DMatrix5(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])

	targetRowIndex := -1

	t, b := 0, rows-1
	for t <= b {
		rowIndex := (t + b) / 2
		row := matrix[rowIndex]
		firstVal, endVal := row[0], row[len(row)-1]

		if target > endVal {
			t = rowIndex + 1
		} else if target < firstVal {
			b = rowIndex - 1
		} else {
			targetRowIndex = rowIndex
			break
		}
	}

	if targetRowIndex == -1 {
		return false
	}

	l, r := 0, cols-1
	for l <= r {
		midIndex := (l + r) / 2
		mid := matrix[targetRowIndex][midIndex]

		if target > mid {
			l = midIndex + 1
		} else if target < mid {
			r = midIndex - 1
		} else {
			return true
		}
	}

	return false
}
