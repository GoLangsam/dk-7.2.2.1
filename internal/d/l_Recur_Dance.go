// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

// ===========================================================================

func (l *L) recurDance(
	d tacher,
	on *On,
	i x.Index,

) {

	{
		d.DoCover(i) // Inline ========================================
	}

	for v := l.optaS[i].Next; v != i; v = l.optaS[v].Next {

		l.CellS[l.Index] = v // Remember cell
		l.Index++            // Incr Level
		d.DoCoverOthers(v)   // Inline ========================================
		on.down()            // Twirl
		d.UnCoverOthers(v)   // Inline ========================================
		l.Index--            // Decr Level
	}

	{
		d.UnCover(i) // Inline ========================================
	}

}

// ===========================================================================
