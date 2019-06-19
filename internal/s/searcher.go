// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/d" // dancing
	"github.com/GoLangsam/dk-7.2.2.1/internal/m" // problem matrix
)

// ========================================================

// Searcher consolidates what's needed to Search.
type Searcher struct {
	m    *m.M // the problem Matrix - intentionally shadowed
	d.D       // the Dancer - can Dance
	pace      // the Pacer  - can Spin or Turn

	// x.Stack

	logOnDone bool
	logOnGoal bool
	logChoice bool
}

// ========================================================

// New returns (a pointer to) a fresh Searcher
// with silent default settings,
// ready to be formed, and
// ready to Search().
func New(M *m.M, setters ...Setter) *Searcher {

	cap := len(M.ItemS)

	a := new(Searcher)
	a.m = M                      // the problem Matrix
	a.D = d.New(M)               // can Dance - implements Dancer
	a.pace.drums = newDrums(cap) // can Count - Beat the Rhythm

	// a.Stack = x.NewStack(cap)

	a.Settings( // defaults
		OnDone(a.onDone),        // YES We have to return
		OnGoal(a.onGoal),        // YES We have a solution
		SetChooser(a.chooseMRV), // YES We have to go on dancing & goaling

		SetPacer(a.pace.Spin),              // ... here we spin & turn
		SetDancer(a.D.DanceFake),           // ... here links only pretend to dance
		SetDancer(a.D.DanceFast),           // ... here links dance faster
		SetDancer(a.D.Dance),               // ... here links dance
		CallOnLeaf(a.pace.drums.leaf.Beat), // Count updates
	)

	a.Settings( // user settings
		setters...)

	return a
}

// ========================================================

// Search is what a Searcher is busy doing.
func (a *Searcher) Search() *Searcher {
	a.D.On.Next()
	return a
}

// ========================================================

// Print prints results of a Search
// depending on the verbosity settings.
//
// Per default: nothing is printed.
func (a *Searcher) Print() *Searcher {
	a.D.Drum.Print()
	a.pace.drums.Print()
	return a
}

// ========================================================
