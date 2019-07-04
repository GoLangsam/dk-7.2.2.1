// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
	"github.com/GoLangsam/do"
)

// ===========================================================================

type chooser func(*x.Item) x.Index

// ===========================================================================

type dancer func(x.Index)

// ===========================================================================

// On consolidates what controls behaviour.
type On struct {
	choose chooser // mandatory - chooses (next) item to cover - use MRV heuristic per default

	search do.It // recursive or flat - algorithm X or C

	next dancer // for recursive dance: Callback Dance
	down do.It  // for recursive dance: Callback Twirl

	Init do.It // optional: Called on Open/Entry/Start/Begin
	Done do.It // optional: Called on Done/Exit/Finished/End

	Skip do.Nok // optional: Called per Enter-Level
	Goal do.It  // optional: Called per Solution
	Fail do.It  // optional: Called per Deadend
	Leaf do.It  // optional: Called per Update
}

// ===========================================================================

func (a On) GetInit() *do.It  { return &a.Init }
func (a On) GetSkip() *do.Nok { return &a.Skip }
func (a On) GetGoal() *do.It  { return &a.Goal }
func (a On) GetFail() *do.It  { return &a.Fail }
func (a On) GetLeaf() *do.It  { return &a.Leaf }
func (a On) GetDone() *do.It  { return &a.Done }

// ===========================================================================
