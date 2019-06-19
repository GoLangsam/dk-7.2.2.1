// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package m

import (
	"fmt"

	"github.com/GoLangsam/dk-7.2.2.1/internal/x"
)

// ===========================================================================

// ShowOption exercice 12
func (a *M) ShowOption(idx x.Index) {

	if idx <= x.Index(len(a.ItemS)) {
		die("ShowOption: used on an item!")
	}

	if idx > x.Index(len(a.OptaS)) {
		die("ShowOption: index too large - no option here!")
	}

	if a.OptaS[idx].Root < 0 { // Spacer
		die("ShowOption: used on a spacer!")
	}

	show := func(i x.Index) {
		fmt.Print(a.NameS[a.OptaS[i].Root], tab)
	}

	a.Do(show).ForEachLineNext(idx)

	top := a.OptaS[idx].Root
	q := a.OptaS[top].Next

	var k int
	for k = 1; q != idx && q != top; k++ {
		q = a.OptaS[q].Next
	}

	name := a.NameS[top]
	if q != top {
		fmt.Print("the option containing", name, " is ", k, " of ", a.OptaS[top].Root, tab)
	} else {
		fmt.Print("the option containing", name, " is not on this list", tab)
	}
	fmt.Println()
}

// ===========================================================================
