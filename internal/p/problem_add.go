// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package p

import (
	"fmt"
	"log"

	"github.com/GoLangsam/dk-7.2.2.1/internal/x"
)

// ===========================================================================

// AddItems returns the extended problem - thus calls may be chained.
//
// Panics iff AddOptions had been called before (with non-empty args).
func (a *P) AddItems(items ...string) *P {
	if len(items) == 0 {
		return a
	}

	if len(a.ItemS) != len(a.OptaS) {
		log.Panic("Cannot add any more items: options have been added already!")
	}

	root := len(a.ItemS) // here the root will be.
	var name string      // empty name for main root.
	a.AddList(root)      // add it to the matrix

	// push things in lockstep
	for _, name = range items {
		a.AddItem(name, root)
	}

	// TODO: consider to set Prev & Next of a.OptaS[root] - currently unused.

	return a
}

// AddOption returns the extended problem - thus calls may be chained.
func (a *P) AddOption(items ...string) *P {
	if len(items) == 0 {
		return a
	}

	if len(a.ItemS) == len(a.OptaS) { // add first trailing spacer
		a.AddMark(-1, 0) // Note: DK starts marking with 0, and decrements. We start negative, using -1.
	}

	c := len(a.OptaS)                        // shall create a.OptaS[c]
	seen := make(map[string]int, len(items)) // to avoid duplicate items in this option.

	for i, name := range items {
		if prev, ok := seen[name]; ok {
			log.Panic(fmt.Sprintf("AddOption: duplicate item `%v`: first seen at %v, now at %v!", name, prev, i))
		}
		seen[name] = i

		a.AddCell(a.MustKnow(x.Name(name))) // append to Column(name-Index)
	}

	a.AddMark(a.OptaS[c-1].Root-1, c)        // add trailing spacer
	a.OptaS[c-1].Next = (c - 1) + len(items) // update preceding spacer

	return a
}

// ===========================================================================
