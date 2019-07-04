// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

// ===========================================================================

func (a tachX) UnCoverOthers(vi x.Index) {
	var (
		hi x.Index // some option adjacent to vi - think: horizontal
		hO *x.Opta // OptaS[di]
		i  x.Index // DoCover(i): i or hO.Root // Change variable in order to ease inlining
	)

	{

		// Beg of UnCoverOthers ==========================================
		hi = vi - 1
		for hi != vi {
			hO = &a.OptaS[hi]
			if hO.Root < 0 { // Spacer
				hi = hO.Next
				continue
			}

			i = hO.Root  // Change variable in order to ease inlining
			a.UnCover(i) // Inline ========================================
			hi--
		}
		// End of UnCoverOthers ==========================================

	}
}

// ===========================================================================

func (a tachX) UnCover(i x.Index) {
	var (
		iI     *x.Item // cover: ItemS[i]
		il, ir x.Index // cover: left, right
		p      x.Index // cover:: DoHide(p)
	)

	{

		// Beg of Cover ======================================================

		iI = &a.ItemS[i]
		//  a.ItemS.ReTach(iI) ===============================================
		il = iI.Prev
		ir = iI.Next
		a.ItemS[il].Next = i
		a.ItemS[ir].Prev = i

		for p = a.OptaS[i].Prev; p != i; p = a.OptaS[p].Prev {
			a.UnHide(p) // Inline ========================================

		}
		// End of Cover ======================================================

	}
}

// ===========================================================================

// UnHide p.Index
func (a tachX) UnHide(p x.Index) {
	var (
		qi     x.Index // hide: index
		qO     *x.Opta // hide: OptaS[qi]
		qu, qd x.Index // hide: up, down
	)

	{
		{

			// Beg of Hide ===================================================
			qi = p - 1
			for qi != p { // ForEachOtherPrev

				qO = &a.OptaS[qi]

				if qO.Root < 0 { // Spacer
					qi = qO.Next
					continue
				}

				//  a.OptaS.ReTach(qq) =======================================
				qu = qO.Prev
				qd = qO.Next
				a.OptaS[qd].Prev = qi
				a.OptaS[qu].Next = qi

				a.OptaS[qO.Root].Root++
				qi--

				// Beg of Update-Count =======================================

				// End of Update-Count =======================================

			}
			// End of Hide ===================================================

		}
	}
}

// ===========================================================================
