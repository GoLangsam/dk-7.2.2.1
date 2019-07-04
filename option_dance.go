// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dl

import (
	"github.com/GoLangsam/do"

	"github.com/GoLangsam/dk-7.2.2.1/internal/d" // dancing
)

// ===========================================================================

/*
// OnDown returns a function which
// sets D.On.Down to f
// and returns it's undo do.Opt.
func OnDown(fs ...func(*d.D)) do.Option {
	return func(any interface{}) do.Opt {
		a := any.(*Searcher)
		return a.DoItSet(&a.D.On.down, fs...)(a)
	}
}
*/

/*
// OnHere returns a function which
// sets D.On.Here to f
// and returns it's undo do.Opt.
func OnHere(f ...func(*d.D)) do.Option {
	return func(any interface{}) do.Opt {
		a := any.(*Searcher)
		return a.DoItSet(&a.D.On.Here, f...)(a) }
	}
}
*/

// ===========================================================================

// OnSkip returns a function which
// sets D.On.Skip to f
// and returns it's undo doFn.
func OnSkip(fs ...func(*d.D) bool) do.Option {
	return func(any interface{}) do.Opt {
		a := aD(any)
		return (&a.On.Skip).Add(a.WrapNok(fs...))(a)
	}
}

// OnGoal returns a function which
// sets D.On.Goal to f
// and returns it's undo do.Opt.
func OnGoal(fs ...func(*d.D)) do.Option {
	return func(any interface{}) do.Opt {
		a := aD(any)
		return (&a.On.Goal).Add(a.WrapIt(fs...))(a)
	}
}

// OnFail returns a function which
// sets D.On.Fail to f
// and returns it's undo do.Opt.
func OnFail(fs ...func(*d.D)) do.Option {
	return func(any interface{}) do.Opt {
		a := aD(any)
		return (&a.On.Fail).Add(a.WrapIt(fs...))(a)
	}
}

// OnLeaf returns a function which
// sets D.On.Leaf to f
// and returns it's undo do.Opt.
func OnLeaf(fs ...func(*d.D)) do.Option {
	return func(any interface{}) do.Opt {
		a := aD(any)
		return (&a.On.Leaf).Add(a.WrapIt(fs...))(a)
	}
}

// ===========================================================================

// VerboseS controls the verbosity setting of package d.
func VerboseS(v bool) do.Option {
	return func(any interface{}) do.Opt {
		prev := verbose
		verbose = v
		return func() do.Opt {
			return VerboseS(prev)(any)
		}
	}
}

// VerboseD controls the verbosity setting of package d.
func VerboseD(v bool) do.Option {
	return func(any interface{}) do.Opt {
		prev := d.GetVerbose()
		d.SetVerbose(v)
		return func() do.Opt {
			return VerboseD(prev)(any)
		}
	}
}

// VerboseM controls the verbosity setting of package m.
func VerboseM(v bool) do.Option {
	return func(any interface{}) do.Opt {
		prev := d.GetVerboseM()
		d.SetVerboseM(v)
		return func() do.Opt {
			return VerboseM(prev)(any)
		}
	}
}

// VerboseX controls the verbosity setting of package x.
func VerboseX(v bool) do.Option {
	return func(any interface{}) do.Opt {
		prev := d.GetVerboseX()
		d.SetVerboseX(v)
		return func() do.Opt {
			return VerboseX(prev)(any)
		}
	}
}

// ===========================================================================
