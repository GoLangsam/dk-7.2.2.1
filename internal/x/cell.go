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

// SetItem stores the item at the current Level.
func (a *Cells) SetItem(main Main, v Index, opta *Opta) {
	a.CellS[a.Level] = cell{main, v, opta}
}

// SetOpta stores the Opta at the current Level.
func (a *Cells) SetOpta(opta *Opta) {
	a.CellS[a.Level].Opta = opta
}

// GetItem returns a structure which represents the item at the current Level.
func (a *Cells) GetItem() (main Main) {
	c := a.CellS[a.Level]
	main = c.Main
	return
}

// GetOpta returns the opta and it's root in OptaS at the current Level.
func (a *Cells) GetOpta() *Opta {
	return a.CellS[a.Level].Opta
}

// GetBoth returns both: the main item and the opta at the current Level.
func (a *Cells) GetBoth() (main Main, opta *Opta) {
	c := a.CellS[a.Level]
	main, opta = c.Main, c.Opta
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
