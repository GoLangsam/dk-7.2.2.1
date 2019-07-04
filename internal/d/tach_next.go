// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

// ===========================================================================

// Next chooses the index of another primary Item
// among the remaining items and returns it for Dance(i).
//
// Chosen is the first item with the smallest number of options.
// (Sometimes called "MRV heuristic", or "S heuristic".)
//
// Next panics iff the list at root is empty
// as no item can be found.
func (a tach) Next(root *x.Item) (here x.Index) {

	Size := x.Index(len(a.OptaS)) // larger than anything we'll find.

	if root.Next == 0 {
		die("Choice called on empty list!")
	}

	var size x.Index
	for curr := root.Next; curr != 0; curr = a.ItemS[curr].Next {
		// TODO: the "non-sharp/sharp preference"-Heuristics
		// if a.NameS[curr] does/doesn't start with `#` {
		//	size = size + len(a.Optas.MarkS) - 1 // #-of-options
		// }

		size = a.OptaS[curr].Root
		if size < Size {
			Size, here = size, curr
		}
	}
	if here == 0 {
		die("Choice found no item!")
	}

	qqq("Chosen:", tab, here)
	return
}

// ========================================================
