// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package m

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x"
)

// ===========================================================================

// M represents the (usually sparse) matrix of some exact cover problem.
//
// Note: The null value is NOT useful! Use `NewMatrix(...)`.
type M struct {
	x.Names // the names
	x.Items // the items
	x.Optas // the options
}

// NewMatrix returns (a pointer to) a new (sparse) matrix M
// initialized with given capacity cap.
func NewMatrix(cap int) *M {
	m := M{
		x.Names{
			make([]x.Name, 0, cap),
			make(map[string]x.Index, cap),
		},
		x.Items{
			make([]x.Item, 0, cap),
			make([]x.Index, 0, cap),
		},
		x.Optas{
			make([]x.Opta, 0, cap*cap), // pre-allocate generously; subsequent Clone will cut the fat
			make([]x.Index, 0, cap),
		},
	}
	return &m
}

// ===========================================================================

// Clone returns a new matrix consisting of a copy of the slice data of a.
//
// Note: The name-dictionary is read-only!
// Only its pointer is copied.
func (a *M) Clone() *M {
	m := M{
		a.Names.Clone(),
		a.Items.Clone(),
		a.Optas.Clone(),
	}

	return &m
}

// ===========================================================================
