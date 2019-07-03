// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package data

import (
	"github.com/GoLangsam/do/id"
)

// NQueens returns a N-Queens problem
// where both rows and ranks are primary.
func NQueens(N int) (lines [][]string) {
	return nQueens(N, false)
}

// NQueensR returns a N-Queens problem
// where 'only' the rows are primary.
func NQueensR(N int) (lines [][]string) {
	return nQueens(N, true)
}

func nQueens(N int, onlyR bool) (lines [][]string) {

	var r = id.S("R-", N)     // Ranks
	var f = id.S("F-", N)     // Files
	var a = id.S("A-", 2*N-1) // DiagA
	var b = id.S("B-", 2*N-1) // DiagB

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

	var ai, bi int // index for Diags

	for fi := range f {
		for ri := range r {
			ai = fi + ri
			bi = N - 1 - fi + ri
			lines = append(lines, []string{r[ri], f[fi], a[ai], b[bi]})
		}
	}
	return
}
