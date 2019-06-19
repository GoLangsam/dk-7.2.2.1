// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x

// ===========================================================================

// ForEachPrev.
func (a Do) ForEachPrev(root int) {

	for curr := a.ItemS[root].Prev; curr != root; curr = a.ItemS[curr].Prev {
		a.Do(curr)
	}
}

// ===========================================================================

// ForEachOptaPrev.
func (a Do) ForEachOptaPrev(root int) {

	for curr := a.OptaS[root].Prev; curr != root; curr = a.OptaS[curr].Prev {
		a.Do(curr)
	}
}

// ===========================================================================

// ForEachLinePrev.
func (a Do) ForEachLinePrev(i int) {

	here := i
	if a.OptaS[here].Root < 0 { // Spacer
		here = a.OptaS[here].Next
	}
	if here == 0 {
		return
	}

	for curr := here; ; {
		a.Do(curr)

		curr--
		if a.OptaS[curr].Root < 0 { // Spacer
			curr = a.OptaS[curr].Next
		}

		if curr == here {
			break
		}
	}
}

// ===========================================================================

// ForEachOtherPrev.
func (a Do) ForEachOtherPrev(i int) {

	if a.OptaS[i].Root < 0 { // Spacer
		a.ForEachLinePrev(i)
	}

	for curr := i - 1; curr != i; {
		if a.OptaS[curr].Root < 0 { // Spacer
			curr = a.OptaS[curr].Next
			continue
		}
		a.Do(curr)
		curr--
	}
}

// ===========================================================================
