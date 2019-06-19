// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package l

// ===========================================================================

// Path represents a search constellation - a (partial) solution.
type Path []Node

// Note: some people use a stack

// AppEnd some Node: aka stack.Push(...)
func (a Path) AppEnd(o Node) Path {
	return append(a, o)
}

// DotDot forgets the last Node: aka stack.Drop(), or Path "..".
func (a Path) DotDot() Path {
	return a[:len(a)-1]
}

// Peek returns the current Path.
//
// The result is intended for quick checks only,
// must NOT be modified
// and is NOT safe for concurrent use!
func (a Path) Peek() Path {
	return a
}

// Path returns (a clone of) the current Path.
//
// Thus, the returned Path is safe for concurrent use.
func (a Path) Path() Path {
	p := make([]Node, len(a))
	copy(p, a)
	return p
}

// ===========================================================================
