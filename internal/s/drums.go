// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x"
)

// "github.com/GoLangsam/container/oneway/drum" // == "github.com/GoLangsam/tees/dance/dancers/drummer/drum"

// drums consolidates all drums (counters) needed.
type drums struct {
	Goal x.Drum // Grooves counts solutions per level
	Fail x.Drum // Deadend counts failures per level
	Call x.Drum // Niveaus counts effort per level
	leaf x.Drum // UpDates counts unLink per level
	// Note: No drum for Push & Pop, as these occur in sync with Call

	Verbose bool
}

// newDrums provides a fresh ensemble of drums
func newDrums(cap int) drums {
	return drums{
		Goal: x.NewDrum("Grooves", cap),
		Fail: x.NewDrum("Deadend", cap),
		Call: x.NewDrum("Niveaus", cap),
		leaf: x.NewDrum("UpDates", cap),
	}
}

// Print has all drums print iff Verbose
func (d *drums) Print() {
	if d.Verbose {
		d.Goal.Print()
		d.Fail.Print()
		d.Call.Print()
		d.leaf.Print()
	}
}
