// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s

// SetDancer returns a function which
// sets pace.On.Next to v
// and returns it's undo doFn.
func SetDancer(v func(int)) Setter {
	return func(a *Searcher) doFn {
		prev := a.pace.On.Next
		a.pace.On.Next = v
		return func() doFn {
			return SetDancer(prev)(a)
		}
	}
}

// OnDone returns a function which
// sets pace.On.Done to v
// and returns it's undo doFn.
func OnDone(v func() bool) Setter {
	return func(a *Searcher) doFn {
		prev := a.pace.On.Done
		a.pace.On.Done = v
		return func() doFn {
			return OnDone(prev)(a)
		}
	}
}

// OnGoal returns a function which
// sets pace.On.Goal to v
// and returns it's undo doFn.
func OnGoal(v func() bool) Setter {
	return func(a *Searcher) doFn {
		prev := a.pace.On.Goal
		a.pace.On.Goal = v
		return func() doFn {
			return OnGoal(prev)(a)
		}
	}
}

// SetChooser returns a function which
// sets pace.On.Fail to v
// and returns it's undo doFn.
func SetChooser(v func() (int, bool)) Setter {
	return func(a *Searcher) doFn {
		prev := a.pace.On.Fail
		a.pace.On.Fail = v
		return func() doFn {
			return SetChooser(prev)(a)
		}
	}
}
