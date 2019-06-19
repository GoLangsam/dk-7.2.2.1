// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s

// VerboseOnSpin returns a function which
// sets pace.verboseOnSpin to v
// and returns it's undo doFn.
func VerboseOnSpin(v bool) Setter {
	return func(a *Searcher) doFn {
		prev := a.pace.beatOnSpin
		a.pace.beatOnSpin = v
		return func() doFn {
			return VerboseOnSpin(prev)(a)
		}
	}
}
