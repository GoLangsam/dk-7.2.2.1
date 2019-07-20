// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x

// ===========================================================================

// Item has a unique id and represents
// a member of a doubly linked list
// implemented as a ring with a sentiel root
// hosted in a slice.
type Item struct {
	Prev, Next Index // pointer-indices
}

// ===========================================================================

// ItemS represents a collection of items.
type ItemS []Item

// ===========================================================================

// Clone returns a new slice consisting of a copy of the data of a.
func (a ItemS) Clone() ItemS {
	c := make([]Item, len(a))
	copy(c, a)
	return c
}

// ===========================================================================

// Items represents a collection of items
// with a list of their marks (= roots).
type Items struct {
	ItemS
	MarkS
}

// ===========================================================================

// Clone returns a new structure consisting of a copy of the data of a.
func (a Items) Clone() Items {
	return Items{
		a.ItemS.Clone(),
		a.MarkS.Clone(),
	}
}

// ===========================================================================

// AppendList appends a root for a new doubly linked list.
func (a ItemS) AppendList(root Index) ItemS {
	c := Index(len(a)) // shall create a[c]

	return append(a, Item{Prev: c, Next: c})
}

// AppendItem appends another item to the doubly linked list rooted at root.
func (a ItemS) AppendItem(root Index) ItemS {
	c := Index(len(a))             // shall create a[c]
	p := c - 1                     // Prev
	a[p].Next, a[root].Prev = c, c // adjust existing Prev.Next & root.Prev

	return append(a, Item{Prev: p, Next: root})
}

// ===========================================================================

// DeTach detaches the given item from the ItemS.
func (a ItemS) DeTach(item Index) {
	a[a[item].Prev].Next = a[item].Next
	a[a[item].Next].Prev = a[item].Prev
}

// ReTach retaches the given item back into the ItemS.
func (a ItemS) ReTach(item Index) {
	a[a[item].Prev].Next = item
	a[a[item].Next].Prev = item
}

// ===========================================================================
// Note: Following methods intend to clarify some semantics.
// For performance reasons they are not called elsewhere but inlined directly.

// Llink returns the left link.
func (a Item) Llink() Index {
	return a.Prev
}

// Rlink returns the right link.
func (a Item) Rlink() Index {
	return a.Next
}

// ===========================================================================
