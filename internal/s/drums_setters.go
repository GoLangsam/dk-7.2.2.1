// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s

// VerboseDrums returns a function which
// sets pace.drums.Verbose to v
// and returns it's undo doFn.
func VerboseDrums(v bool) Setter {
	return func(a *Searcher) doFn {
		prev := a.pace.drums.Verbose
		a.pace.drums.Verbose = v
		return func() doFn {
			return VerboseDrums(prev)(a)
		}
	}
}

// VerboseDrumsGoal returns a function which
// sets pace.drums.Goal.UseMap to v
// and returns it's undo doFn.
func VerboseDrumsGoal(v bool) Setter {
	return func(a *Searcher) doFn {
		prev := a.pace.drums.Goal.UseMap
		a.pace.drums.Goal.UseMap = v
		return func() doFn {
			return VerboseDrumsGoal(prev)(a)
		}
	}
}

// VerboseDrumsFail returns a function which
// sets pace.drums.Fail.UseMap to v
// and returns it's undo doFn.
func VerboseDrumsFail(v bool) Setter {
	return func(a *Searcher) doFn {
		prev := a.pace.drums.Fail.UseMap
		a.pace.drums.Fail.UseMap = v
		return func() doFn {
			return VerboseDrumsFail(prev)(a)
		}
	}
}

// VerboseDrumsCall returns a function which
// sets pace.drums.Call.UseMap to v
// and returns it's undo doFn.
func VerboseDrumsCall(v bool) Setter {
	return func(a *Searcher) doFn {
		prev := a.pace.drums.Call.UseMap
		a.pace.drums.Call.UseMap = v
		return func() doFn {
			return VerboseDrumsCall(prev)(a)
		}
	}
}

// VerboseDrumsLeaf returns a function which
// sets pace.drums.leaf.UseMap to v
// and returns it's undo doFn.
func VerboseDrumsLeaf(v bool) Setter {
	return func(a *Searcher) doFn {
		prev := a.pace.drums.leaf.UseMap
		a.pace.drums.leaf.UseMap = v
		return func() doFn {
			return VerboseDrumsLeaf(prev)(a)
		}
	}
}
