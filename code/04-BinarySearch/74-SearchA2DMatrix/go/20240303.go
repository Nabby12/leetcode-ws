package main

// https://leetcode.com/problems/search-a-2d-matrix/description/
func searchA2DMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])

	targetRow := -1
	top, bottom := 0, rows-1
	for top <= bottom {
		row := (top + bottom) / 2
		if target > matrix[row][len(matrix[row])-1] {
			top = row + 1
		} else if target < matrix[row][0] {
			bottom = row - 1
		} else {
			targetRow = row
			break
		}
	}

	if targetRow == -1 {
		return false
	}

	targetRows := matrix[targetRow]
	l, r := 0, cols-1
	for l <= r {
		mid := (l + r) / 2
		midVal := targetRows[mid]
		if midVal == target {
			return true
		} else if target > midVal {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
