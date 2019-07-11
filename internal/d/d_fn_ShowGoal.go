// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"strings"
)

// ===========================================================================
/*
// Do wraps matrix data, local backtracking data
// (and some optional functions).
//
// Usage: a.Do(f).SomeThing(withThis)
type Do struct {
	*m.M
	*x.Cells
	Do []func(*D)
}

// ===========================================================================

// Do returns a Do (which wraps search data and the optional functions) and
// which provides several useful methods visiting an ongoing search.
func (a *D) Do(fs ...func(*D)) Do {
	return Do{&a.M, &a.L.Cells, append([]func(*D){}, fs...)}
}
*/
// ===========================================================================

// LogGoal sends a titled "Solution:" to the current logger.
func LogGoal(a *D) {
	see(a.ShowGoal("Solution:"))
}

// ShowGoal returns a human-readable generic presentation of a solution
// preceded by a title, iff some is given.
func (a *D) ShowGoal(title string) string {
	var b strings.Builder
	a.WriteTitle(&b, title)
	for _, c := range a.Cells.Range() {
		a.M.WriteOptionLine(&b, c.Index)
	}
	return b.String()
}

// ===========================================================================

// WriteTitle writes a title line (iff not empty)
// preceded by D.Name and M.Name
// into the provided strings.Builder.
func (a *D) WriteTitle(b *strings.Builder, title string) {
	if title != "" {
		spc := " "
		b.WriteString(string(a.Name))
		b.WriteString(spc)
		b.WriteString(string(a.M.Name))
		b.WriteString(spc)
		b.WriteString(title)
		b.WriteString(eol)
	}
}

// ===========================================================================

// ShowOption returns a generic human-readable presentation
// of all options found in the current backtrack buffer.
func (a *D) ShowOption() string {
	var b strings.Builder
	for _, c := range a.Cells.Range() {
		b.WriteString(a.M.ShowOption(c.Index))
	}
	return b.String()
}

// ===========================================================================

// Permille returns an estimate of the completion of a search
// based on the results of OptionCellIsKofN of each option cell
// found in the current backtrack buffer (as of formula (27))
// as truncated integer permille (%o) value.
//
// Note: As of now, OptionCellIsKofN is neither stored nor memoized
// and thus is a little costly - use with care (if at all)
// when performance is a priority.
func (a *D) Permille() int {
	var psum float64
	var prod float64 = 1

	for _, c := range a.Cells.Range() {
		k, n := a.OptionCellIsKofN(c.Index)
		prod *= float64(n)
		psum += float64(k-1) / prod
	}

	psum += 1 / (2 * prod)
	return int(psum * 1000)
}

// ===========================================================================
