// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/do"

	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

// ===========================================================================

type chooser func(*x.Item) x.Main

// ===========================================================================

type dancer func(x.Main)

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

// GetInit returns (a pointer to) the function On.Init.
func (a On) GetInit() *do.It { return &a.Init }

// GetSkip returns (a pointer to) the function On.Skip.
func (a On) GetSkip() *do.Nok { return &a.Skip }

// GetGoal returns (a pointer to) the function On.Goal.
func (a On) GetGoal() *do.It { return &a.Goal }

// GetFail returns (a pointer to) the function On.Fail.
func (a On) GetFail() *do.It { return &a.Fail }

// GetLeaf returns (a pointer to) the function On.Leaf.
func (a On) GetLeaf() *do.It { return &a.Leaf }

// GetDone returns (a pointer to) the function On.Done.
func (a On) GetDone() *do.It { return &a.Done }

// ===========================================================================
