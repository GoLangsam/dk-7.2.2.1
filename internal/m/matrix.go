// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package m

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

// ===========================================================================

// M represents the (usually sparse) matrix of some exact cover problem.
//
// Note: The null value is NOT useful!
type M struct {
	lines   [][]string
	x.Name  // the name of the problem
	x.Names // the names
	x.Items // the items
	x.Optas // the options
}

// ===========================================================================

// Clone returns a new matrix consisting of a copy of the slice data of a.
//
// Note: The name-dictionary is read-only!
// Only its pointer is copied.
func (a *M) Clone() *M {
	m := M{
		a.lines,
		a.Name,
		a.Names.Clone(),
		a.Items.Clone(),
		a.Optas.Clone(),
	}

	return &m
}

// ===========================================================================
