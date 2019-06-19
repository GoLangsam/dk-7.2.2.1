// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

// ===========================================================================

func (a *D) DanceFake(here int) {
	fake := here
	a.ItemS.DeTach(fake)
	a.Stack.Push(here)
	if a.On.Leaf != nil {
		a.On.Leaf(len(a.Stack))
	}
	a.On.Next()
	a.Stack.Drop()
	a.ItemS.ReTach(fake)
}

// ===========================================================================
