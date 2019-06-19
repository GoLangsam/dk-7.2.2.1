// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s

// SetPacer returns a function which
// sets D.On.Next to v
// and returns it's undo doFn.
func SetPacer(v func()) Setter {
	return func(a *Searcher) doFn {
		prev := a.D.On.Next
		a.D.On.Next = v
		return func() doFn {
			return SetPacer(prev)(a)
		}
	}
}

// CallOnHere returns a function which
// sets D.On.Here to v
// and returns it's undo doFn.
func CallOnHere(v func() int) Setter {
	return func(a *Searcher) doFn {
		prev := a.D.On.Here
		a.D.On.Here = v
		return func() doFn {
			return CallOnHere(prev)(a)
		}
	}
}

// CallOnLeaf returns a function which
// sets D.On.Leaf to v
// and returns it's undo doFn.
func CallOnLeaf(v func(int)) Setter {
	return func(a *Searcher) doFn {
		prev := a.D.On.Leaf
		a.D.On.Leaf = v
		return func() doFn {
			return CallOnLeaf(prev)(a)
		}
	}
}

// CountUpdatesPerCell returns a function which
// sets D.Drum.UseMap to v
// and returns it's undo doFn.
func CountUpdatesPerCell(v bool) Setter {
	return func(a *Searcher) doFn {
		prev := a.D.Drum.UseMap
		a.D.Drum.UseMap = v
		return func() doFn {
			return CountUpdatesPerCell(prev)(a)
		}
	}
}
