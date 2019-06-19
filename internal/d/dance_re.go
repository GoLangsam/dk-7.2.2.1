// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

// ===========================================================================

func (a *D) UnCoverOthers(vi Index) {
	var hi Index // some option adjacent to vi - think: horizontal
	var hO *Opta // OptaS[di]
	var i Index  // UnCover(i): i or hO.Root // Change variable in order to ease inlining

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

func (a *D) UnCover(i Index) {
	var iI *Item     // cover: ItemS[i]
	var il, ir Index // cover: left, right
	var p Index      // cover: UnHide(p)

	{
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
}

// ===========================================================================

// UnHide p.Index
func (a *D) UnHide(p Index) {
	var qi Index     // hide: index
	var qO *Opta     // hide: OptaS[qi]
	var qu, qd Index // hide: up, down

	{
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
					a.OptaS[qu].Next = qi
					a.OptaS[qd].Prev = qi

					a.OptaS[qO.Root].Root++
					qi--
				}
				// End of Hide ===================================================

			}
		}
	}
}

// ===========================================================================
