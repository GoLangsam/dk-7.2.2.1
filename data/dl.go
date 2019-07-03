// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package data

import (
	"github.com/GoLangsam/dk-7.2.2.1"
)

func SmallMatrixM() dl.Matrix {
	return dl.Problem(SmallMatrix()...).Matrix()
}

// NQueensM returns a N Queens problem matrix.
// where both rows and ranks are primary.
func NQueensM(N int) dl.Matrix {
	return dl.Problem(nQueens(N, false)...).Matrix()
}

// NQueensRM returns a N-Queens problem matrix
// where 'only' the rows are primary.
func NQueensRM(N int) dl.Matrix {
	return dl.Problem(nQueens(N, true)...).Matrix()

}
