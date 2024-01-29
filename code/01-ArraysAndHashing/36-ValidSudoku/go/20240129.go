package main

import "strconv"

// https://leetcode.com/problems/valid-sudoku/description/
func validSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	grids := [9][9]bool{}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			num, err := strconv.Atoi(string(board[row][col]))
			if err != nil {
				continue
			}
			num--
			gridIndex := col/3 + (row/3)*3
			if rows[row][num] || cols[col][num] || grids[gridIndex][num] {
				return false
			}
			rows[row][num] = true
			cols[col][num] = true
			grids[gridIndex][num] = true
		}
	}

	return true
}
