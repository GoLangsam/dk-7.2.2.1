// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package m

import (
	"fmt"
	"strings"

	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
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

// OptionCellIsKofN returns for the given i its position k
// among the total on n members of its option line,
// or panics iff i is not cell of an option.
func (a *M) OptionCellIsKofN(i x.Index) (k, n int) {
	a.AssertOptionCell(i)

	ri := a.OptaS[i].Root
	rO := &a.OptaS[ri]

	n = int(rO.Root)

	q := rO.Next
	for k = 1; q != i && q != ri; k++ {
		q = a.OptaS[q].Next
	}
	return
}

// ===========================================================================

// ShowOption exercice 12
func (a *M) ShowOption(i x.Index) string {
	var b strings.Builder

	a.WriteOptionLine(&b, i)
	b.WriteString(string(a.NameS[a.OptaS[i].Root]))
	b.WriteString(tab)
	k, n := a.OptionCellIsKofN(i) // includes AssertOptionCell
	b.WriteString(fmt.Sprint("=", k, "/", n))
	b.WriteString(eol)
	return b.String()
}

// ===========================================================================

// WriteOptionLine writes the option line of given v
// into the provided strings.Builder.
func (a *M) WriteOptionLine(b *strings.Builder, v x.Index) {
	tab := func() { b.WriteString(tab) }
	show := func(i x.Index) { b.WriteString(string(a.NameS[a.OptaS[i].Root])); tab() }

	for h := v; ; {
		if a.OptaS[h].Root < 0 { // Spacer
			h = a.OptaS[h].Prev
			continue
		}
		show(h)
		h++
		if h != v {
			break
		}
	}
	b.WriteString(eol)
}

// ===========================================================================
