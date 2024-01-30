package main

import "strconv"

// https://leetcode.com/problems/valid-sudoku/description/
func validSudoku2(board [][]byte) bool {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	grids := [3][3][9]bool{}

	for r := 0; r < len(rows); r++ {
		for c := 0; c < len(cols); c++ {
			num, err := strconv.Atoi(string(board[r][c]))
			if err != nil {
				continue
			}

			num--

			if rows[r][num] || cols[c][num] || grids[r/3][c/3][num] {
				return false
			}

			rows[r][num] = true
			cols[c][num] = true
			grids[r/3][c/3][num] = true
		}
	}

	return true
}
