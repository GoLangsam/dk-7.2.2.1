// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package m

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x"
)

// ===========================================================================

// AddList appends a root for a new itemlist rooted at `root`.
func (a *M) AddList(root int) {

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
func (a *M) AddItem(name string, root int) {

	a.Names.Dict.LearnOnce(x.Name(name), len(a.ItemS))
	a.Names.NameS = a.Names.NameS.AppendName(name)

	a.Items.ItemS = a.Items.ItemS.AppendItem(root)

	a.Optas.OptaS = a.Optas.OptaS.AppendRoot(root) // keep optas in sync
}

// AddMark appends a spacer. Note: The mark should be < 0.
func (a *M) AddMark(mark, prev int) {

	a.Optas.MarkS = a.Optas.MarkS.AppendMark(len(a.OptaS))
	a.Optas.OptaS = a.Optas.OptaS.AppendMark(mark, prev)
}

// AddCell appends a cell to the (vertical) list rooted at item `root`.
func (a *M) AddCell(root int) {

	a.Optas.OptaS = a.Optas.OptaS.AppendOpta(root)
}

// ===========================================================================
