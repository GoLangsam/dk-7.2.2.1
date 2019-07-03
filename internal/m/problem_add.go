// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package m

import (
	"fmt"

	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

// ===========================================================================

func getCap(lines ...[]string) (capI, capO int) {
	var isOption bool

	for _, line := range lines {
		size := len(line)
		if size == 0 {
			isOption = true
			continue
		}

		if isOption {
			capO += size + 1 // one more for root
		} else {
			capI += size + 1 // one more for root
		}

	}
	capO += capI // + Items
	capO += 1    // + trailing spacer
	return
}

// ===========================================================================

// AddLines returns the extended problem - thus calls may be chained.
func (a *p) addLines(lines ...[]string) *p {
	var isOption bool

	for _, line := range lines {
		if len(line) == 0 {
			isOption = true
			continue
		}

		if isOption {
			a.AddOption(line...)
		} else {
			a.AddItems(line...)
		}
	}

	return a
}

// ===========================================================================

// AddItems returns the extended problem - thus calls may be chained.
//
// Panics iff AddOptions had been called before (with non-empty args).
func (a *p) AddItems(items ...string) *p {
	if len(items) == 0 {
		return a
	}

	if len(a.ItemS) != len(a.OptaS) {
		die("Cannot add any more items: options have been added already!")
	}

	root := x.Index(len(a.ItemS)) // here the root will be.
	var name string               // empty name for main root.
	a.AddList(root)               // add it to the matrix

	// push things in lockstep
	for _, name = range items {
		a.AddItem(name, root)
	}

	// TODO: consider to set Prev & Next of a.OptaS[root] - currently unused.

	return a
}

// AddOption returns the extended problem - thus calls may be chained.
func (a *p) AddOption(items ...string) *p {
	if len(items) == 0 {
		return a
	}

	if len(a.ItemS) == len(a.OptaS) { // add first trailing spacer
		a.AddMark(-1, 0) // Note: DK starts marking with 0, and decrements. We start negative, using -1.
	}

	c := x.Index(len(a.OptaS))               // shall create a.OptaS[c]
	seen := make(map[string]int, len(items)) // to avoid duplicate items in this option.

	for i, name := range items {
		if prev, ok := seen[name]; ok {
			die(fmt.Sprintf("AddOption: duplicate item `%v`: first seen at %v, now at %v!", name, prev, i))
		}
		seen[name] = i

		a.AddCell(a.MustKnow(x.Name(name))) // append to Column(name-Index)
	}

	a.AddMark(a.OptaS[c-1].Root-1, c)                 // add trailing spacer
	a.OptaS[c-1].Next = (c - 1) + x.Index(len(items)) // update preceding spacer

	return a
}

// ===========================================================================
