// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x

// ===========================================================================

// Main represents an item/column of a problem matrix.
//
// Invariant: Item.Index == Opta.Index
type Main = Index

// ===========================================================================

// Cell represents an element of a problem matrix.
//
// Invariant: Main.Item == OptaS[Opta.Root]
type cell struct {
	Main  // the main root: Item & Opta
	Index // the position of the opta
	*Opta // the opta
}

// Cells is a buffer for visited cells.
type Cells struct {
	Level Index  // the level - the depth of recursion
	CellS []cell // the cells visited per level - compose the solution
}

// ===========================================================================

// NewCells returns a fresh cell-buffer of capacity cap.
func NewCells(cap int) Cells {
	return Cells{CellS: make([]cell, cap)}
}

// ===========================================================================

// SetCell stores the details of the cell at the current Level.
func (a *Cells) SetCell(main Main, v Index, opta *Opta) {
	a.CellS[a.Level] = cell{main, v, opta}
}

// GetCell returns the details of the cell at the current Level.
func (a *Cells) GetCell() (main Main, v Index, opta *Opta) {
	c := a.CellS[a.Level]
	main, v, opta = c.Main, c.Index, c.Opta
	return
}

// ===========================================================================

// Range returns the currently active CellS
func (a *Cells) Range() []cell {
	return a.CellS[:a.Level]
}

// ===========================================================================

// Clone returns a new slice consisting of a copy of the data of a.
func (a *Cells) Clone() []cell {
	c := make([]cell, a.Level)
	copy(c, a.CellS)
	return c
}

// ===========================================================================
