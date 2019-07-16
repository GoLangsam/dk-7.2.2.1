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
	i x.Main,
) {

	{
		d.DoCover(i) // Inline ========================================
	}

	vO := &l.optaS[i]
	v := vO.Next
	for v != i {

		vO = &l.optaS[v]    // read cell
		l.SetCell(i, v, vO) // Remember cell
		l.Level++           // Incr Level
		d.DoCoverOthers(v)  // Inline ========================================
		on.down()           // Twirl
		d.UnCoverOthers(v)  // Inline ========================================
		l.Level--           // Decr Level
		v = vO.Next         // next cell pointer
	}

	{
		d.UnCover(i) // Inline ========================================
	}

}

// ===========================================================================
