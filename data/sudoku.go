// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package data

import (
	"github.com/GoLangsam/do/id"
)

/*
Sodoku:

Square (r,c)
Rows: r 1...9
Cols: c 1...9
Blok: b 1...9

k = Digit

---
Then every sudoku square S = (s(ij)) corresponds naturally to the solution of a master exact cover problem
whose 9 * 9 * 9 = 729 options are
`	p(ij)	r(ik)	c(jk)	b(x(ij)k)' for 0 < i,j < 10, 0 < k < 10, and x = 3*(i/3) + j/3 (30)

and whose 4 * 9 * 9 = 324 items are
	p(ij)	r(ik)	c(jk)	b(x(ij)k)

The reason is that option (30) is chosen with parameters (i, j, k ) if and only if
	s(ij) = k

Item p(ij) must be covered by exactly one of the nine options that fill cell s(ij)
item r(ik) must be covered by exactly one of the nine options that put k in row i...
item b(xk) must be covered by exactly one of the nine options that put k in box x.

Got it?
*/

func sudokuBlock(i, j int) int { return (i/3)*3 + j/3 }

// Sudoku returns a Sudoku problem for the given hint.
func Sudoku(name string, hint SudokuHint) (string, [][]string) {

	const N = 9 // Size of sudoku square
	const D = 9 // No of different digits in sudoku square

	plcs := id.S("P-", N) // Place
	rows := id.S("R-", N) // Rows
	cols := id.S("C-", N) // Cols
	blks := id.S("B-", N) // Blocks

	p := [][]string{}
	r := [][]string{}
	c := [][]string{}
	b := [][]string{}

	for _, plc := range plcs {
		p = append(p, id.S(plc+"-", N))
	}

	for _, row := range rows {
		r = append(r, id.S(row+"V=", D))
	}

	for _, col := range cols {
		c = append(c, id.S(col+"V=", D))
	}

	for _, blk := range blks {
		b = append(b, id.S(blk+"V=", D))
	}

	itemS := []string{}
	optaS := [][]string{}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if hint[i][j] == 0 {
				itemS = append(itemS,
					p[i][j],
				)
			}
			if !hint.IsDinR(j, i) { // k=j, i=i
				itemS = append(itemS,
					r[i][j],
				)
			}
			if !hint.IsDinC(j, i) { // k=j, j=i
				itemS = append(itemS,
					c[i][j],
				)
			}
			if !hint.IsDinB(j, i) { // k=j, x=i
				itemS = append(itemS,
					b[i][j],
				)
			}
			for k := range hint.Hint(i, j) {
				x := sudokuBlock(i, j)
				if hint[i][j] == 0 && !hint.IsDinR(k, i) && !hint.IsDinC(k, j) && !hint.IsDinB(k, x) {
					optaS = append(optaS, []string{
						p[i][j], r[i][k], c[j][k], b[x][k],
					})
				}
			}
		}
	}

	return name, append([][]string{itemS, {}}, optaS...) // items & separator & options
}
