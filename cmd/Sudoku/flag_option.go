// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/GoLangsam/do"

	"github.com/GoLangsam/dk-7.2.2.1"
	"github.com/GoLangsam/dk-7.2.2.1/internal/d"
)

// ========================================================

func getFlagOption() do.Option {
	// v, vg, rh, vd, vb, vc bool // as in tees/cmd/NQueensR
	// d.Dancer = dancing.New(vb, rh, false)

	options := []do.Option{}

	if *times {
		options = append(options, dl.LogTime())
	}
	if *sizes {
		options = append(options, dl.LogSize())
	}
	if *goals {
		options = append(options, dl.OnGoal(d.LogGoal))
	}
	if *choos {
		options = append(options, dl.VerboseD(*choos))
	}
	if *verbs {
		options = append(options, dl.VerboseM(*verbs), dl.VerboseX(*verbs))
	}
	if *dGoal {
		options = append(options, dl.CountGrooves(*dGoalMap))
	} // Count solutions
	if *dFail {
		options = append(options, dl.CountDeadEnd(*dFailMap))
	} // Count dead-ends
	if *dCall {
		options = append(options, dl.CountNiveaus(*dCallMap))
	} // Count effort
	if *dLeaf {
		options = append(options, dl.CountUpDates(*dLeafMap))
	} // Count updates
	if *dleaf {
		options = append(options, dl.CountUpdates(*dleafMap))
	} // Count updates per cell

	if *progress {
		options = append(options, dl.LogProgress(*progressCnt))
	}

	return do.OptionJoin(options...)
}

// ========================================================
