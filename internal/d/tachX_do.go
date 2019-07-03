// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

// ===========================================================================

func (a tachX) DoCoverOthers(vi x.Index) {
	var (
		hi x.Index // some option adjacent to vi - think: horizontal
		hO *x.Opta // OptaS[di]
		i  x.Index // DoCover(i): i or hO.Root // Change variable in order to ease inlining
	)

	{

		// Beg of DoCoverOthers ==========================================
		hi = vi + 1
		for hi != vi {

			hO = &a.OptaS[hi]

			if hO.Root < 0 { // Spacer
				hi = hO.Prev
				continue
			}

			i = hO.Root  // Change variable in order to ease inlining
			a.DoCover(i) // Inline ========================================

			hi++
		}
		// End of DoCoverOthers ==========================================

	}
}

// ===========================================================================

func (a tachX) DoCover(i x.Index) {
	var (
		iI     *x.Item // cover: ItemS[i]
		il, ir x.Index // cover: left, right
		p      x.Index // cover:: DoHide(p)
	)

	{
		{

			// Beg of Cover ======================================================
			for p = a.OptaS[i].Next; p != i; p = a.OptaS[p].Next {

				a.DoHide(p) // Inline ========================================

			}

			iI = &a.ItemS[i]
			//  a.ItemS.DeTach(iI) ===============================================
			il = iI.Prev
			ir = iI.Next
			a.ItemS[il].Next = ir
			a.ItemS[ir].Prev = il
			// End of Cover ======================================================

		}
	}
}

// ===========================================================================

// DoHide p.Index
func (a tachX) DoHide(p x.Index) {
	var (
		qi     x.Index // hide: index
		qO     *x.Opta // hide: OptaS[qi]
		qu, qd x.Index // hide: up, down
	)

	{
		{
			{

				// Beg of Hide ===================================================
				qi = p + 1
				for qi != p { // ForEachOtherNext

					qO = &a.OptaS[qi]

					if qO.Root < 0 { // Spacer
						qi = qO.Prev
						continue
					}

					//  a.OptaS.DeTach(qO) =======================================
					qu = qO.Prev
					qd = qO.Next
					a.OptaS[qd].Prev = qu
					a.OptaS[qu].Next = qd

					a.OptaS[qO.Root].Root--
					qi++

					// Beg of Update-Count =======================================
					if a.Drum != nil {
						a.Drum.Beat(int(qi)) // count update per Opta
					}
					if a.It != nil && *a.It != nil {
						(*a.It).Do() // count updates
					}
					// End of Update-Count =======================================
				}
				// End of Hide ===================================================

			}
		}
	}
}

// ===========================================================================
