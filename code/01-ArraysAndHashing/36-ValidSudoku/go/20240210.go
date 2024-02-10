package main

import (
	"strconv"
)

// https://leetcode.com/problems/valid-sudoku/description/
func validSudoku3(board [][]byte) bool {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	grids := [3][3][9]bool{}

	for i, _ := range rows {
		for j, _ := range cols {
			num, err := strconv.Atoi(string(board[i][j]))
			if err != nil {
				continue
			}

			num--

			gridRow, gridCol := i/3, j/3
			if rows[i][num] || cols[j][num] || grids[gridRow][gridCol][num] {
				return false
			}

			rows[i][num] = true
			cols[j][num] = true
			grids[gridRow][gridCol][num] = true
		}
	}

	return true
}
