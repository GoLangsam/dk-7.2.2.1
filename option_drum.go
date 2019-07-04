// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dl

import (
	"github.com/GoLangsam/container/oneway/drum"
	"github.com/GoLangsam/dk-7.2.2.1/internal/d" // dancing
	"github.com/GoLangsam/do"
)

// ===========================================================================

func beatShow(a *d.D, name string, cap int, useMap bool) (beat, show do.It) {
	drum := drum.NewDrum(name, cap)
	drum.UseMap = useMap

	beat = a.WrapIt(func(d *d.D) { drum.Beat(int(d.L.Index)) })
	show = drum.Print
	return
}

// CountNiveaus counts effort
func CountNiveaus(perLevel bool) do.Option {
	return func(any interface{}) do.Opt {
		a := aD(any)
		beat, show := beatShow(a, "Niveaus", len(a.ItemS), perLevel)
		undoDrum := (&a.On.Skip).Add(do.NokIt(beat))(a)
		undoDone := (&a.On.Done).Add(show)(a)

		return func() do.Opt {
			return do.OptJoin(undoDrum, undoDone)
		}

	}
}

// CountGrooves counts solutions
func CountGrooves(perLevel bool) do.Option {
	return func(any interface{}) do.Opt {
		a := aD(any)
		beat, show := beatShow(a, "Grooves", len(a.ItemS), perLevel)
		undoDrum := (&a.On.Goal).Add(beat)(a)
		undoDone := (&a.On.Done).Add(show)(a)

		return func() do.Opt {
			return do.OptJoin(undoDrum, undoDone)
		}

	}
}

// CountDeadEnd counts failures
func CountDeadEnd(perLevel bool) do.Option {
	return func(any interface{}) do.Opt {
		a := aD(any)
		beat, show := beatShow(a, "DeadEnd", len(a.ItemS), perLevel)
		undoDrum := (&a.On.Fail).Add(beat)(a)
		undoDone := (&a.On.Done).Add(show)(a)

		return func() do.Opt {
			return do.OptJoin(undoDrum, undoDone)
		}

	}
}

// CountUpDates counts unLink
func CountUpDates(perLevel bool) do.Option {
	return func(any interface{}) do.Opt {
		a := aD(any)
		beat, show := beatShow(a, "UpDates", len(a.ItemS), perLevel)
		undoDrum := (&a.On.Leaf).Add(beat)(a)
		undoDone := (&a.On.Done).Add(show)(a)

		return func() do.Opt {
			return do.OptJoin(undoDrum, undoDone)
		}

	}
}

// ===========================================================================

// CountUpdates returns a function which
// sets D.Drum.UseMap to perCell
// and returns it's undo do.Opt.
func CountUpdates(perCells bool) do.Option {
	return func(any interface{}) do.Opt {
		a := aD(any)

		undoDrum := countUpdates(perCells)(a)
		undoDone := (&a.On.Done).Add(a.Drum.Print)(a) // TODO: this panics iff we have no a.Drum!

		return func() do.Opt {
			return do.OptJoin(undoDrum, undoDone)
		}
	}
}

func countUpdates(perCells bool) do.Option {
	return func(any interface{}) do.Opt {
		a := aD(any)
		var prev bool
		if a.Drum != nil {
			prev = a.Drum.UseMap
			a.Drum.UseMap = perCells
		}
		return func() do.Opt {
			return countUpdates(prev)(a)
		}
	}
}

// ===========================================================================
