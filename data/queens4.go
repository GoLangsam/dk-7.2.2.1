// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package data

import (
	"github.com/GoLangsam/dk-7.2.2.1"
)

/*
The 4 queens problem is just the task of covering
	eight primary columns (R0,R1,R2,R3, F0, F1, F2, F3)
		corresponding to 4 ranks + 4 files,
	while using at most one element in each of the secondary columns
	(A0,A1,A2,A3,A4,A5,A6, B0,B1,B2,B3,B4,B5,B6) corresponding to diagonals, given the sixteen rows
*/

// FourQueens returns a handmade 4-Queen matrix.
func FourQueens() dl.Matrix {

	return dl.Problem([][]string{
		[]string{"R0", "R1", "R2", "R3", "F0", "F1", "F2", "F3"}, // Ranks & Files
		[]string{"A0", "A1", "A2", "A3", "A4", "A5", "A6"},       // Diagonals
		[]string{"B0", "B1", "B2", "B3", "B4", "B5", "B6"},       // Diagonals backward
		[]string{},
		[]string{"R0", "F0", "A0", "B3"},
		[]string{"R0", "F1", "A1", "B4"},
		[]string{"R0", "F2", "A2", "B5"},
		[]string{"R0", "F3", "A3", "B6"},

		[]string{"R1", "F0", "A1", "B2"},
		[]string{"R1", "F1", "A2", "B3"},
		[]string{"R1", "F2", "A3", "B4"},
		[]string{"R1", "F3", "A4", "B5"},

		[]string{"R2", "F0", "A2", "B1"},
		[]string{"R2", "F1", "A3", "B2"},
		[]string{"R2", "F2", "A4", "B3"},
		[]string{"R2", "F3", "A5", "B4"},

		[]string{"R3", "F0", "A3", "B0"},
		[]string{"R3", "F1", "A4", "B1"},
		[]string{"R3", "F2", "A5", "B2"},
		[]string{"R3", "F3", "A6", "B3"},
	}...).Matrix()

}
