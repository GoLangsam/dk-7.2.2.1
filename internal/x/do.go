// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x

// ===========================================================================

// Do wraps matrix data and a Do function and provides iterators:
//  ForEach....      - iterates the items (in ItemS) given a root
//  ForEachOpta....  - iterates the items (in OptaS) given a root
//  ForEachLine....  - iterates an option line
//  ForEachOther.... - OptaS: any other option of the option line
//
//  .... = Next iterates forward,
//  .... = Prev iterates backward,
//
// Usage: a.Do(f).ForEach....(startHere)
type Do struct {
	*Items
	*Optas
	Stack []Index
	Do    func(Index)
}

// ===========================================================================
