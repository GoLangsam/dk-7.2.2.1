// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/container/oneway/drum"

	"github.com/GoLangsam/dk-7.2.2.1/internal/m" // problem matrix
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

// ===========================================================================

func SetVerboseM(v bool) {
	m.SetVerbose(v)
}

func GetVerboseM() bool {
	return m.GetVerbose()
}

func SetVerboseX(v bool) {
	x.SetVerbose(v)
}

func GetVerboseX() bool {
	return x.GetVerbose()
}

// ===========================================================================

// Cell represents a non-zero element of the matrix.
type cell = x.Index

// CellS buffers visited cells for backtracking.
type CellS []cell

// ===========================================================================

// L consolidates local data for backtracking
// and may be visited by On actions.
type L struct {
	x.Index // the level - the depth of recursion
	CellS   // the cell visited per level - compose the solution
	// Note: cell's *Item is useful to present some solution!
	optaS x.OptaS // the options
}

// ===========================================================================

type D struct {
	x.Name               // the name of the alorithm in use
	m.M                  // the problem Matrix
	*L                   // the local variables of the Dancer - for visitors
	*drum.Drum           // the update Drum - incremented in DoHide
	*On                  // the plug-in functions: what to do when
	tacher               // can do and reCover
	curr       []x.Index // partial solution to begin with
}

// ===========================================================================

// New returns a fresh Dancer
// based on the given problem Matrix.
func newD(M func() m.M, useKind, useDrum bool) *D {
	a := D{M: M(), On: &On{}}
	if useDrum {
		a.Drum = drum.NewDrum("UpOptaS", len(a.M.OptaS))
	}

	a.curr = []x.Index{}

	if useKind {
		tach := tachX{tach{a.ItemS, a.OptaS, a.Drum, &a.On.Leaf}}
		a.tacher = tach
		a.choose = tach.Next
	} else {
		tach := tachX{tach{a.ItemS, a.OptaS, a.Drum, &a.On.Leaf}}
		a.tacher = tach
		a.choose = tach.Next
	}

	a.L = &L{CellS: make([]cell, len(a.M.ItemS)), // sure this is large enough
		optaS: a.OptaS,
	}

	return &a
}

// ===========================================================================

// Search is what a Searcher is busy doing
// after OnInit and before OnDone.
func (a *D) Search(from ...x.Index) {
	a.On.Init.Do()
	a.SearchFrom(from)
	a.On.Done.Do()
}

// ===========================================================================

// SearchFrom is what a Searcher is busy doing
// after winding down to from and
// before rewinding back from from
// so he may try again from elsewhere.
func (a *D) SearchFrom(from []x.Index) {
	a.curr = from
	//a.doWind()
	a.On.search.Do()
	//a.reWind()
	a.curr = []x.Index{}
}

// ===========================================================================
