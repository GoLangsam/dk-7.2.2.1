// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"strings"
)

// ===========================================================================

// LogGoal sends a titled "Solution:" to the current logger.
func LogGoal(a *D) {
	see(a.ShowGoal("Solution:"))
}

// ShowGoal returns a human-readable generic presentation of a solution
// preceded by a title, iff some is given.
func (a *D) ShowGoal(title string) string {
	var b strings.Builder
	a.WriteTitleLine(&b, title)
	for _, c := range a.Cells.Range() {
		a.M.WriteOptionLine(&b, c.Index)
		b.WriteString(eol)
	}
	return b.String()
}

// ===========================================================================

// WriteTitleLine writes a title line (iff not empty)
// preceded by D.Name and M.Name
// into the provided strings.Builder.
func (a *D) WriteTitleLine(b *strings.Builder, title string) {
	if title != "" {
		a.WriteTitle(b, title)
		b.WriteString(eol)
	}
}

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
	}
}

// ===========================================================================

// ShowOption returns a generic human-readable presentation
// of all options found in the current backtrack buffer.
func (a *D) ShowOption() string {
	var b strings.Builder
	for _, c := range a.Cells.Range() {
		a.M.WriteOption(&b, c.Index)
		b.WriteString(eol)
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
