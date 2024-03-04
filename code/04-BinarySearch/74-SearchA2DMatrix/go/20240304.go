package main

// https://leetcode.com/problems/search-a-2d-matrix/description/
func searchA2DMatrix2(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])

	targetRow := -1
	t, b := 0, rows-1
	for t <= b {
		row := (t + b) / 2
		if target > matrix[row][len(matrix[row])-1] {
			t = row + 1
		} else if target < matrix[row][0] {
			b = row - 1
		} else {
			targetRow = row
			break
		}
	}

	if targetRow == -1 {
		return false
	}

	targetRowMatrix := matrix[targetRow]
	l, r := 0, cols-1
	for l <= r {
		mid := (l + r) / 2
		midVal := targetRowMatrix[mid]
		if target == midVal {
			return true
		} else if target > midVal {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
