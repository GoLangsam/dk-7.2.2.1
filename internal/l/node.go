// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package l

import (
	"log"
)

// ===========================================================================

// Node represents the behaviour of any element of a doubly linked list structure.
type Node interface {
	Nxt() int
	Prv() int
	Top() int
	SetN(int)
	SetP(int)
	SetR(int)
}

// ===========================================================================

// BiLi represents the bilateral links of a doubly linked list structure.
type BiLi struct {
	Prev, Next int
}

func (a BiLi) Nxt() int { return a.Next }
func (a BiLi) Prv() int { return a.Prev }
func (a BiLi) Top() int { log.Panic("I do not know about root!"); return 0 }

func (a BiLi) SetN(link int) { a.Next = link }
func (a BiLi) SetP(link int) { a.Prev = link }
func (a BiLi) SetR(link int) { log.Panic("I do not know about root!") }

// ===========================================================================

// SpMa represents an element of a sparse matrix
// implemented as a doubly doubly linked list structure
// hosted in some array (slice).
type SpMa struct {
	Root int
	BiLi
}

func (a SpMa) Top() int      { return a.Root }
func (a SpMa) SetR(link int) { a.Root = link }

// ===========================================================================
