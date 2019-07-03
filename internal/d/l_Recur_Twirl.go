// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

// ===========================================================================

// CallBack from Dance - as D.On.Next

// recurTwirl is what a dancer does: keep dancing or not
func (l *L) recurTwirl(
	d tacher,
	on *On,
	OptaS x.OptaS,
	root *x.Item,
	next chooser,
) {

	if on.Skip != nil && on.Skip.Do() { // => X8 - skip level
		return
	}
	if OptaS[0].Next == 0 { // => X8 (all items have been covered),
		if on.Goal != nil {
			on.Goal.Do()
		} // we have a Solution!
		return // >>>>>>>>>>
	}
	i := next(root)
	if OptaS[i].Next == i { // => X7 - deadend
		if on.Fail != nil {
			on.Fail.Do()
		} // we have a Dead end!
		return // >>>>>>>>>>
	}

	on.next(i) // Dance
}

// ===========================================================================
