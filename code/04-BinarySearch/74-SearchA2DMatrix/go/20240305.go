package main

// https://leetcode.com/problems/search-a-2d-matrix/description/
func searchA2DMatrix3(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])

	targetRowIndex := -1
	r, b := 0, rows-1
	for r <= b {
		midRowIndex := (r + b) / 2
		midRow := matrix[midRowIndex]

		firstVal, endVal := midRow[0], midRow[len(midRow)-1]

		if firstVal <= target && target <= endVal {
			targetRowIndex = midRowIndex
			break
		}

		if target > firstVal {
			r = midRowIndex + 1
		} else {
			b = midRowIndex - 1
		}
	}

	if targetRowIndex == -1 {
		return false
	}

	targetRow := matrix[targetRowIndex]
	l, r := 0, cols-1
	for l <= r {
		mid := (l + r) / 2
		midVal := targetRow[mid]

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
