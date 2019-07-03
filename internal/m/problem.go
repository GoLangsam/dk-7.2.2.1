// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package m

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

// ===========================================================================

// P represents a problem under construction.
//
// Note: The null value is NOT useful!
type p struct {
	*M // (a pointer to) the matrix.
}

// ===========================================================================

// Problem returns a new problem matrix
// populated with the given lines.
//
// Options (if any) must be preceded by an empty line.
//
// The first line defines the primary items,
// subsequent non-empty lines (if any) define (groups of) secondary items.
//
// The problem is ready to accept secondary items (if any)
// and further options
// by subsequent calls to its methods AddItems(...) resp. AddOption(...).
//
// It panics iff some duplicate item name is encountered,
// or iff some item has an empty string as its name
// or iff some option refers to an item name more than once
// or iff there are no items.
func Problem(name string, lines ...[]string) *p {
	capI, capO := getCap(lines...)
	if capI == 0 {
		die("Problem: need some items!")
	}

	a := p{&M{
		lines,
		x.Name(name),
		x.Names{make([]x.Name, 0, capI), make(map[x.Name]x.Index, capI)},
		x.Items{make([]x.Item, 0, capI), []x.Index{}},
		x.Optas{make([]x.Opta, 0, capO), []x.Index{}},
	}} // make new problem matrix

	return (&a).addLines(lines...)
}

// ===========================================================================

// Items returns a new problem matrix
// populated with the given primary items and
// ready to accept secondary items (if any) and options
// by subsequent calls to its methods AddItems(...) resp. AddOption(...).
//
// It panics iff some duplicate item name is encountered,
// or iff some item has an empty string as its name.
func Items(name string, items ...string) *p {
	N := len(items)
	if N < 1 {
		die("Items: need one item - at least!")
	}

	capI := N + 2             // allocate two more: for primary and secondary roots
	capO := capI * capI
	a := p{&M{
		[][]string{items},
		x.Name(name),
		x.Names{make([]x.Name, 0, capI), make(map[x.Name]x.Index, capI)},
		x.Items{make([]x.Item, 0, capI), []x.Index{}},
		x.Optas{make([]x.Opta, 0, capO), []x.Index{}},
	}} // make new problem matrix

	a.AddItems(items...)

	return &a
}

// ===========================================================================

// Matrix returns (a pointer to) the matrix.
func (a *p) Matrix() *M {
	return a.M
}

// ===========================================================================
