// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package m

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

// ===========================================================================

// AddList appends a root for a new itemlist rooted at `root`.
func (a *M) addList(root x.Index) {

	var name string

	if len(a.NameS) == 0 { // only for first root - name is intentionally empty
		a.Names.Dict.LearnOnce(x.Name(name), root)
	} else {
		name = tab
	}

	a.Names.NameS = a.Names.NameS.AppendName(name)

	a.Items.MarkS = a.Items.MarkS.AppendMark(root)
	a.Items.ItemS = a.Items.ItemS.AppendList(root)

	a.Optas.OptaS = a.Optas.OptaS.AppendNull() // keep optas in sync
}

// AddItem appends a named item to the itemlist rooted at `root`.
func (a *M) addItem(name string, root x.Index) {

	{
		a.Names.Dict.LearnOnce(x.Name(name), x.Index(len(a.ItemS)))
	}

	a.Names.NameS = a.Names.NameS.AppendName(name)

	a.Items.ItemS = a.Items.ItemS.AppendItem(root)

	a.Optas.OptaS = a.Optas.OptaS.AppendRoot(root) // keep optas in sync
}

// AddMark appends a spacer. Note: The mark should be < 0.
func (a *M) addMark(mark, prev x.Index) {

	c := x.Index(len(a.OptaS)) // shall create a[c]

	a.Optas.MarkS = a.Optas.MarkS.AppendMark(c)
	a.Optas.OptaS = a.Optas.OptaS.AppendMark(mark, prev)
}

// AddCell appends a cell to the (vertical) list rooted at item `root`.
func (a *M) addCell(root x.Index) {

	a.Optas.OptaS = a.Optas.OptaS.AppendOpta(root)
}

// ===========================================================================
