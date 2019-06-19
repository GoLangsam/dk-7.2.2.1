// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x"
)

// ===========================================================================

func (a *D) DoCoverOthers(vi Index) {
	var hi Index // some option adjacent to vi - think: horizontal
	var hO *Opta // OptaS[di]
	var i Index  // DoCover(i): i or hO.Root // Change variable in order to ease inlining

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

func (a *D) DoCover(i Index) {
	var iI *Item     // cover: ItemS[i]
	var il, ir Index // cover: left, right
	var p Index      // cover:: DoHide(p)

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
func (a *D) DoHide(p Index) {
	var qi Index     // hide: index
	var qO *Opta     // hide: OptaS[qi]
	var qu, qd Index // hide: up, down

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
					a.OptaS[qu].Next = qd
					a.OptaS[qd].Prev = qu

					a.OptaS[qO.Root].Root--
					qi++

					// Beg of Update-Count =======================================
					a.Drum.Cnt++ // count update per Opta
					if a.Drum.UseMap {
						a.Drum.Map[qi]++
					}
					if a.On.Leaf != nil { // count update per Level
						a.On.Leaf(x.Index(len(a.Stack)))
					}
					// End of Update-Count =======================================
				}
				// End of Hide ===================================================

			}
		}
	}
}

// ===========================================================================
