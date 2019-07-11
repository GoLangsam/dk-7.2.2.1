// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/do"

	"github.com/GoLangsam/dk-7.2.2.1/internal/m" // problem matrix
)

// ===========================================================================

// AlgorithmX returns (a pointer to) a fresh Searcher
// with silent default settings.
func AlgorithmX(M func() m.M, setters ...do.Option) D {
	a := algorithmX(M, false, true)
	a.Settings(setters...) // user settings

	return a
}

// AlgorithmC returns (a pointer to) a fresh Searcher
// with silent default settings.
func AlgorithmC(M func() m.M, setters ...do.Option) D {
	a := algorithmX(M, true, true)
	a.Settings(setters...) // user settings

	return a
}

// ===========================================================================

func algorithmX(M func() m.M, useKind, useDrum bool) D {
	a := newD(M, useKind, useDrum)
	if useKind {
		a.Name = "AlgorithmC"
	} else {
		a.Name = "AlgorithmX"
	}

	a.On.search = func() {
		a.L.algorithmX(a.tacher, a.On, &a.ItemS[0], a.On.choose)
	}

	return a
}

// ===========================================================================
