// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package data

func Hint29a() (string, SudokuHint) {
	return "Sudoku29a",
		[][]int{
			{0, 0, 3, 0, 1, 0, 0, 0, 0},
			{4, 1, 5, 0, 0, 0, 0, 9, 0},
			{2, 0, 6, 5, 0, 0, 3, 0, 0},
			{5, 0, 0, 0, 8, 0, 0, 0, 9},
			{0, 7, 0, 9, 0, 0, 0, 3, 2},
			{0, 3, 8, 0, 0, 4, 0, 6, 0},
			{0, 0, 0, 2, 6, 0, 4, 0, 3},
			{0, 0, 0, 3, 0, 0, 0, 0, 8},
			{3, 2, 0, 0, 0, 7, 9, 5, 0},
		}
}

func Hint29b() (string, SudokuHint) {
	return "Sudoku29b",
		[][]int{
			{0, 0, 0, 0, 0, 0, 3, 0, 0},
			{1, 0, 0, 4, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 0, 5},
			{9, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 2, 6, 0, 0},
			{0, 0, 0, 0, 5, 3, 0, 0, 0},
			{0, 5, 0, 8, 0, 0, 0, 0, 0},
			{0, 0, 0, 9, 0, 0, 0, 7, 0},
			{0, 8, 3, 0, 0, 0, 0, 4, 0},
		}
}

func Hint29c() (string, SudokuHint) {
	return "Sudoku29c",
		[][]int{
			{0, 3, 0, 0, 1, 0, 0, 0, 0},
			{0, 0, 0, 4, 0, 0, 1, 0, 0},
			{0, 5, 0, 0, 0, 0, 0, 9, 0},
			{2, 0, 0, 0, 0, 0, 6, 0, 4},
			{0, 0, 0, 0, 3, 5, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0},
			{4, 0, 0, 6, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 5, 0},
			{0, 9, 0, 0, 0, 0, 0, 0, 0},
		}
}

// SudokuHint represents a partially filled
// Sudoku grid awaiting to be solved.
type SudokuHint [][]int

var all = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// Hint returns the digits applicable to cell i,j
// which is either a predefined digit, or all.
func (a SudokuHint) Hint(i, j int) []int {
	if a[i][j] == 0 {
		return all
	}
	return []int{a[i][j]}
}

func (a SudokuHint) IsDinR(k, i int) bool {
	d := k + 1
	for _, ki := range a[i] {
		if ki == d {
			return true
		}
	}
	return false
}

func (a SudokuHint) IsDinC(k, j int) bool {
	d := k + 1
	for _, r := range a {
		if r[j] == d {
			return true
		}
	}
	return false
}

func (a SudokuHint) IsDinB(k, b int) bool {
	d := k + 1
	x := sudokuBlock
	for hi, r := range a {
		for hj, ki := range r {
			if ki == d && x(hi, hj) == b {
				return true
			}
		}
	}
	return false
}
