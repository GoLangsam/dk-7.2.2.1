// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x

// ===========================================================================

// MarkS represents a collection of markers, where such mark
// is either a root of (primary or some secondary) items
// or a spacer around some option-line.
type MarkS []Index

// ===========================================================================

// Clone returns a new slice consisting of a copy of the data of a.
func (a MarkS) Clone() MarkS {
	c := make([]Index, len(a))
	copy(c, a)
	return c
}

// ===========================================================================

// AppendMark appends a mark.
func (a MarkS) AppendMark(mark Index) MarkS {

	return append(a, mark)
}

// ===========================================================================
