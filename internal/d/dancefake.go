// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x"
)

// ===========================================================================

func (a *D) DanceFake(here x.Index) {
	fake := here
	a.ItemS.DeTach(fake)
	a.Stack.Push(here)
	if a.On.Leaf != nil {
		a.On.Leaf(x.Index(len(a.Stack)))
	}
	a.On.Next()
	a.Stack.Drop()
	a.ItemS.ReTach(fake)
}

// ===========================================================================
