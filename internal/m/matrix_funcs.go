// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package m

import (
	"fmt"

	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all You need
)

// ===========================================================================

// AssertPrimaryItem panics iff i is not a primary item.
func (a *M) AssertPrimaryItem(i x.Index) {

	if a.OptaS[i].Root < 0 { // Spacer
		die("PrimaryItem: used on a spacer!")
	}

	idx := int(i)
	if idx > len(a.ItemS) {
		die("PrimaryItem: used on an item!")
	}

}

// AssertOptionCell panics iff i is not an option cell.
func (a *M) AssertOptionCell(i x.Index) {

	if a.OptaS[i].Root < 0 { // Spacer
		die("OptionCell: used on a spacer!")
	}

	idx := int(i)
	if idx <= len(a.ItemS) {
		die("OptionCell: used on an item!")
	}

	if idx > len(a.OptaS) {
		die("OptionCell: index too large - no option here!")
	}

}

// ===========================================================================

func (a *M) OptionCellIsKofN(i x.Index) (k, n int) {
	a.AssertOptionCell(i)

	ri := a.OptaS[i].Root
	rO := &a.OptaS[ri]

	n = int(rO.Root)

	q := rO.Next
	for k = 1; q != i && q != ri; k++ {
		q = a.OptaS[q].Next
	}
	if q == ri {
		k, n = 0, 0
	}
	return
}

// ===========================================================================

// ShowOption exercice 12
func (a *M) ShowOption(i x.Index) {
	k, n := a.OptionCellIsKofN(i) // includes AssertOptionCell

	show := func(i x.Index) { fmt.Print(a.NameS[a.OptaS[i].Root], tab) }
	a.Do(show).ForEachLineNext(i)

	name := a.NameS[a.OptaS[i].Root]
	if k != 0 {
		fmt.Print("the option containing", name, " is ", k, " of ", n, tab)
	} else {
		fmt.Print("the option containing", name, " is not on this list", tab)
	}
	fmt.Println()
}

// ===========================================================================
