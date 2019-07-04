// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/m" // problem matrix
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

// ===========================================================================

func AlgorithmX(M func() m.M, useKind, useDrum bool) *D {
	a := newD(M, useKind, useDrum)

	mainS := make([]x.Index, len(a.M.ItemS)) // sure this is large enough

	a.On.search = func() {
		a.L.algorithmX(a.tacher, a.On, &a.ItemS[0], a.On.choose, mainS)
	}

	return a
}

// ===========================================================================
