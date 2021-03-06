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

// SetVerboseM sets the verbosity of package m to v.
func SetVerboseM(v bool) {
	m.SetVerbose(v)
}

// GetVerboseM returns the verbosity of package m.
func GetVerboseM() bool {
	return m.GetVerbose()
}

// SetVerboseX sets the verbosity of package x to v.
func SetVerboseX(v bool) {
	x.SetVerbose(v)
}

// GetVerboseX returns the verbosity of package x.
func GetVerboseX() bool {
	return x.GetVerbose()
}

// ===========================================================================

// L consolidates local backtracking data
// and may be visited by On actions.
type L struct {
	x.Cells         // the buffer of cells visited per level - compose the solution
	optaS   x.OptaS // the options
}

// ===========================================================================

// D consolidates the data of a dancer.
type D struct {
	x.Name               // the name of the alorithm in use
	m.M                  // the problem Matrix
	*L                   // the local backtracking data - for visitors
	*drum.Drum           // the update Drum - incremented in DoHide
	*On                  // the plug-in functions: what to do when
	tacher               // can do and reCover
	curr       []x.Index // partial solution to begin with
}

// ===========================================================================

// New returns a fresh Dancer
// based on the given problem Matrix.
func newD(M func() m.M, useKind, useDrum bool) D {
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

	a.L = &L{x.NewCells(len(a.M.ItemS)), a.OptaS} // sure this is large enough

	return a
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
