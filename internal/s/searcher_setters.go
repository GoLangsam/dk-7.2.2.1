// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s

// LogOnDone returns a function which
// sets logOnLeaf to v
// and returns it's undo doFn.
func LogOnDone(v bool) Setter {
	return func(a *Searcher) doFn {
		prev := a.logOnDone
		a.logOnDone = v
		return func() doFn {
			return LogOnDone(prev)(a)
		}
	}
}

// LogOnGoal returns a function which
// sets logOnGoal to v
// and returns it's undo doFn.
func LogOnGoal(v bool) Setter {
	return func(a *Searcher) doFn {
		prev := a.logOnGoal
		a.logOnGoal = v
		return func() doFn {
			return LogOnGoal(prev)(a)
		}
	}
}

// LogChoice returns a function which
// sets logChoice to v
// and returns it's undo doFn.
func LogChoice(v bool) Setter {
	return func(a *Searcher) doFn {
		prev := a.logChoice
		a.logChoice = v
		return func() doFn {
			return LogChoice(prev)(a)
		}
	}
}
