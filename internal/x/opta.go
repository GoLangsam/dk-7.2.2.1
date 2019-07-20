// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x

// ===========================================================================

// Opta represents an option - somthing one can 'optare' (=choose) from.
type Opta struct {
	Root Index
	Item
}

// ===========================================================================

// OptaS represents a collection of optas.
type OptaS []Opta

// ===========================================================================

// Clone returns a new slice consisting of a copy of the data of a.
func (a OptaS) Clone() OptaS {
	c := make([]Opta, len(a))
	copy(c, a)
	return c
}

// ===========================================================================

// Optas represents a collection of optas
// with a list of their marks (= spacers).
type Optas struct {
	OptaS
	MarkS
}

// ===========================================================================

// Clone returns a new structure consisting of a copy of the data of a.
func (a Optas) Clone() Optas {
	return Optas{
		a.OptaS.Clone(),
		a.MarkS.Clone(),
	}
}

// ===========================================================================

// AppendRoot appends another option root.
func (a OptaS) AppendRoot(root Index) OptaS {
	c := Index(len(a)) // shall create a[c]
	a[root].Root--     // decr length in Root-Item[root]

	return append(a, Opta{Item: Item{Prev: c, Next: c}})
}

// AppendOpta appends another option to the list for the item rooted at root.
func (a OptaS) AppendOpta(root Index) OptaS {
	c := Index(len(a))             // shall create a[c]
	p := a[root].Prev              // Prev
	a[p].Next, a[root].Prev = c, c // adjust existing Prev.Next & Root.Prev
	a[root].Root++                 // incr length in Root-Item[root]

	return append(a, Opta{Item: Item{Prev: p, Next: root}, Root: root})
}

// AppendNull appends another empty option.
func (a OptaS) AppendNull() OptaS {

	return append(a, Opta{})
}

// AppendMark appends a spacer. The mark must be < 0.
func (a OptaS) AppendMark(mark, prev Index) OptaS {
	if !(mark < 0) {
		die("mark must be negative!")
	}

	return append(a, Opta{Item: Item{Prev: prev}, Root: mark})
}

// ===========================================================================

// DeTach detaches the given item from the OptaS
// and decrements their count.
func (a OptaS) DeTach(item Index) {
	a[a[item].Prev].Next = a[item].Next
	a[a[item].Next].Prev = a[item].Prev
	a[a[item].Root].Root--
}

// ReTach retaches the given item back into the OptaS
// and increments their count.
func (a OptaS) ReTach(item Index) {
	a[a[item].Prev].Next = item
	a[a[item].Next].Prev = item
	a[a[item].Root].Root++
}

// ===========================================================================
// Note: Following methods intend to clarify some semantics.
// For performance reasons they are not called elsewhere but inlined directly.

// Top returns the header.
func (a Opta) Top() Index {
	return a.Root
}

// Ulink returns the up link.
func (a Opta) Ulink() Index {
	return a.Prev
}

// Dlink returns the down link.
func (a Opta) Dlink() Index {
	return a.Next
}

// Col returns the column index.
//
// Applicable to OptaS[i] for i > len(ItemS)..
func (a Opta) Col() Index {
	return a.Root
}

// Len returns the length of a column.
//
// Applicable to OptaS[i] for 0 <= i <= len(ItemS).
func (a Opta) Len() int {
	return int(a.Root)
}

// Up returns the up link.
func (a Opta) Up() Index {
	return a.Prev
}

// Down returns the down link.
func (a Opta) Down() Index {
	return a.Next
}

// ===========================================================================
