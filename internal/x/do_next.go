// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x

// ===========================================================================

// ForEachNext.
func (a Do) ForEachNext(root int) {

	for curr := a.ItemS[root].Next; curr != root; curr = a.ItemS[curr].Next {
		a.Do(curr)
	}
}

// ===========================================================================

// ForEachOptaNext.
func (a Do) ForEachOptaNext(root int) {

	for curr := a.OptaS[root].Next; curr != root; curr = a.OptaS[curr].Next {
		a.Do(curr)
	}
}

// ===========================================================================

// ForEachLineNext.
func (a Do) ForEachLineNext(i int) {

	here := i
	if a.OptaS[here].Root < 0 { // Spacer
		here = a.OptaS[here].Prev
	}
	if here == 0 {
		return
	}

	for curr := here; ; {
		a.Do(curr)

		curr++
		if a.OptaS[curr].Root < 0 { // Spacer
			curr = a.OptaS[curr].Prev
		}

		if curr == here {
			break
		}
	}
}

// ===========================================================================

// ForEachOtherNext.
func (a Do) ForEachOtherNext(i int) {

	if a.OptaS[i].Root < 0 { // Spacer
		a.ForEachLineNext(i)
	}

	for curr := i + 1; curr != i; {
		if a.OptaS[curr].Root < 0 { // Spacer
			curr = a.OptaS[curr].Prev
			continue
		}
		a.Do(curr)
		curr++
	}
}

// ===========================================================================
