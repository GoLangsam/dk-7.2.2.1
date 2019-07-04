// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/m" // problem matrix
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

// ===========================================================================

func RecursiveX(M func() m.M, useKind, useDrum bool) *D {
	a := newD(M, useKind, useDrum)
	if useKind {
		a.Name = "RecursiveC"
	} else {
		a.Name = "RecursiveX"
	}

	a.On.next = func(m x.Index) {
		a.L.recurDance(a.tacher, a.On, m)
	} // for recursive version
	a.On.down = func() {
		a.L.recurTwirl(a.tacher, a.On, &a.ItemS[0], a.On.choose)
	} // for recursive version

	a.On.search = a.On.down

	return a
}

// ===========================================================================
