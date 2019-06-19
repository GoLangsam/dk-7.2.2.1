// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x

// ===========================================================================

// Name is a unique id and represents an item.
type Name string

// ===========================================================================

// NameS represents a collection of names.
type NameS []Name

// ===========================================================================

// Names represents a collection of names with a lookup dictionary.
type Names struct {
	NameS
	Dict
}

// ===========================================================================

// Clone returns a new slice consisting of a copy of the data of a.
func (a NameS) Clone() NameS {
	c := make([]Name, len(a))
	copy(c, a)
	return c
}

// ===========================================================================

// Clone returns a new slice consisting of a copy of the data of a.
func (a Names) Clone() Names {
	c := Names{
		make([]Name, len(a.NameS)),
		a.Dict,
	}
	copy(c.NameS, a.NameS)
	return c
}

// ===========================================================================

// AppendName appends another name to a.
func (a NameS) AppendName(name string) NameS {
	return append(a, Name(name))
}

// ===========================================================================
