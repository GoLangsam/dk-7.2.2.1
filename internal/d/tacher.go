// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

// ===========================================================================

// Tacher represents someone who knows
// how to DeTach and ReTach nodes of the matrix.
type tacher interface {
	DoCover(i x.Main)
	UnCover(i x.Main)
	DoCoverOthers(v x.Index)
	UnCoverOthers(v x.Index)
}

// SlowTacher represents someone who knows how to hide.
type slowTacher interface {
	tacher

	DoHide(p x.Index)
	UnHide(p x.Index)
}

// ===========================================================================
