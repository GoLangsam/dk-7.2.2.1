// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package data

import (
	"github.com/GoLangsam/dk-7.2.2.1"
)

// NQueensR returns 'only' the rows as primary.
func NQueensR(N int) dl.Matrix {
	return dl.Problem(nQueens(N, true)...).Matrix()
}

// NQueens returns the N Queens problem matrix.
func NQueens(N int) dl.Matrix {
	return dl.Problem(nQueens(N, false)...).Matrix()
}

// nQueens returns 'only' the rows
func nQueens(N int, onlyR bool) (lines [][]string) {

	var r = IDs("R-", N)     // Ranks
	var f = IDs("F-", N)     // Files
	var a = IDs("A-", 2*N-1) // DiagA
	var b = IDs("B-", 2*N-1) // DiagB

	if onlyR {
		lines = append(lines, r)
		lines = append(lines, f)
		lines = append(lines, a)
		lines = append(lines, b)
	} else {
		lines = append(lines, append(r, f...))
		lines = append(lines, append(a, b...))
	}

	lines = append(lines, []string{})

	for fi := range f {
		for ri := range r {
			lines = append(lines, []string{r[ri], f[fi], a[fi+ri], b[N-1-fi+ri]})
		}
	}
	return
}
