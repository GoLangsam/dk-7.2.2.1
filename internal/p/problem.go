// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package p allows to build a problem and to obtain its matrix representation.
package p

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/m"
)

// ===========================================================================

// P represents a problem under construction.
//
// Note: The null value is NOT useful!
type P struct {
	*m.M // (a pointer to) the matrix.
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
func Problem(lines ...[]string) *P {
	var cap int

	for _, line := range lines {
		size := len(line)
		if size == 0 {
			break
		}
		cap += size + 1 // one more for root
	}

	if cap == 0 {
		die("Problem: need some items!")
	}

	a := P{m.NewMatrix(cap)} // make new problem matrix
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
func Items(items ...string) *P {
	N := len(items)
	if N < 1 {
		die("Items: need one item - at least!")
	}

	cap := N + 2             // allocate two more: for primary and secondary roots
	a := P{m.NewMatrix(cap)} // make new problem matrix

	a.AddItems(items...)

	return &a
}

// ===========================================================================

// Matrix returns (a pointer to) the matrix.
func (a *P) Matrix() *m.M {
	return a.M
}

// ===========================================================================
