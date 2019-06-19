// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x

/* ===========================================================================
Some notes about structure and invariants of some matrix `a` built with `items`:

	N := len(items)      // # of items
	I := len(a.ItemS)    // len
	Z := len(a.OptaS) -1 // index of last spacer

	I == N + 1
	N == I - 1

	N == a.ItemS[0].root
	N == -a.OptaS[0].root // DK leaves a.OptaS[0] alone; we build in tandem

	a.ItemS[0].next == index of first Item iff != 0
	a.ItemS[0].prev == index of last Item iff != 0

	a.OptaS[I] is the first spacer. Thus:
	a.OptaS[I].root == -1

	a.OptaS[Z] is the last spacer. Thus:
	a.OptaS[Z].Root == -(#-option-lines + 1)

	// traverse primary items:
	for i := a.ItemS[0].next; i != 0; i = a.ItemS[i].next {
		a.OptaS[i] represents 'the-column-root':
		a.OptaS[i].root == len(options-with-i)
		a.OptaS[i].next == index of first option-with-i
		a.OptaS[i].prev == index of last  option-with-i

	  	// traverse options:
		for x := a.OptaS[i].next; x != i; x = a.OptaS[x].next {
			a.OptaS[x] represents an option with Item[i]
		}
	}

*/

// ===========================================================================
