// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/m"
	"github.com/GoLangsam/dk-7.2.2.1/internal/x"
)

// ===========================================================================

type D struct {
	Items // represents the items
	Optas // represents the item-headers and the (spaced) options
	Index // represents the Level, the depth of recursion
	Stack // represents the Constellation
	Drum  // represents the Updatecounter
	On    // represents the plug-in functions: what to do when
}

// New returns a fresh Dancer
// based on data cloned from the given problem Matrix.
func New(M *m.M) D {
	d := D{
		Items: M.Items,
		Optas: M.Optas,
		Stack: x.NewStack(len(M.ItemS)), // len(M.ItemS) == -M.OptaS[0].Root,
		Drum:  x.NewDrum("UpOptaS", len(M.OptaS)),
	} // can dance
	d.On.Here = d.Choice
	return d
}

// ===========================================================================

// Choice returns the index of a primary Item - its Root is index for Dance(i).
//
// It panics iff the list rooted at ItemS[0] is empty.
// TODO: DK Version for AlgX - put elsewhere !!!
func (a *D) Choice() (here Index) {

	Size := Index(len(a.OptaS)) // larger than anything we'll find.
	root := &a.ItemS[0]

	if root.Next == 0 {
		panic("Choice called on empty list!")
	}

	var size Index
	for curr := root.Next; curr != 0; curr = a.ItemS[curr].Next {
		size = a.OptaS[curr].Root
		// TODO: the "non-sharp/sharp preference"-Heuristics
		// if a.NameS[curr] does/doesn't start with `#` {
		//	size = size + len(a.Optas.MarkS) - 1 // #-of-options
		// }
		/*
			if size == 0 {
				here, found = -1, false
				break
			}
		*/
		if size < Size {
			Size = size
			here = curr
			//	found = true
		}
	}
	/*
		if a.logChoice { // ... we may show it
			fmt.Println("Chosen:", tab, here, tab, found, tab)
		}
	*/
	return
}
