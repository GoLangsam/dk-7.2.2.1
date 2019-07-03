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
	OptaS x.OptaS,
	c x.Index,

	) {

	var ( 
		v x.Index // some option under c - think: vertical
	)

	{
		{
			d.DoCover(c) // Inline ========================================
		}
	}

	for v = OptaS[c].Next; v != c; v = OptaS[v].Next {

		l.CellS[l.Index] = v // Remember cell
		l.Index++            // Incr Level
		d.DoCoverOthers(v)   // Inline ========================================
		on.down()            // Twirl
		d.UnCoverOthers(v)   // Inline ========================================
		l.Index--            // Decr Level


	}

	{
		{
			d.UnCover(c) // Inline ========================================
		}
	}

}

// ===========================================================================
