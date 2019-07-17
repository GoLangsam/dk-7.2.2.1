// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
)

// ========================================================

var (
	/*
		beg = flag.Int("Beg", 10, "Begin with Boardsize (min=1, max=End)")
		end = flag.Int("End", 12, "End with Boardsize (min=1)")
	*/

	/*
		fastF = flag.Bool("goFast", false, "Use DanceFast, or")
		listF = flag.Bool("goList", false, "Use list/Danc, or")
		slowF = flag.Bool("goSlow", false, "Use DanceSlow")

		sortF = flag.Bool("doSort", false, "Use Sort-by-size")
	*/

	/*
		beats = flag.Bool("b", false, "Beats : Print counters per Level (sets -d)")
		drums = flag.Bool("d", false, "Drums : Print counters")
		rhyth = flag.Bool("r", false, "Rhythm: Collect counters")
	*/

	goals = flag.Bool("s", false, "Verbose Scoring: Print solutions")
	times = flag.Bool("t", false, "Log time on done")
	choos = flag.Bool("u", false, "Log all choices made")
	verbs = flag.Bool("v", false, "Verbose: may log a lot of details")

	progress    = flag.Bool("p", false, "Log progress")
	progressCnt = flag.Int64("pc", 10, "any N-million updates")

	/*
		tests = flag.Bool("x", false, "Execute Chooser-test")
		dancs = flag.Bool("y", false, "Execute Dancing-test (-Fast -norm -Slow -List)")
	*/

	sizes = flag.Bool("z", false, "Log sizes on init")

	dGoal = flag.Bool("cg", false, "Count Grooves")
	dFail = flag.Bool("cd", false, "Count DeadEnd")
	dCall = flag.Bool("cn", false, "Count Niveaus")
	dLeaf = flag.Bool("cu", false, "Count UpDates")
	dleaf = flag.Bool("cl", false, "Count UpDates")

	dGoalMap = flag.Bool("cg+", false, "Count Grooves per Level")
	dFailMap = flag.Bool("cd+", false, "Count DeadEnd per Level")
	dCallMap = flag.Bool("cn+", false, "Count Niveaus per Level")
	dLeafMap = flag.Bool("cu+", false, "Count UpDates per Level")
	dleafMap = flag.Bool("cl+", false, "Count Updates per Cells")

//	s = flag.String("s", " ", "separator")
)

func flagParse() {
	flag.Parse()
	/*
		if *beats {
			*drums = true
		}
		if *drums {
			*rhyth = true
		}
	*/

	/*
		if *beg < 1 {
			*beg = 1
		}
		if *end < 1 {
			*end = 1
		}
		if *beg > *end {
			*beg = *end
		}
	*/

	/*
		if *slowF && *fastF {
			die("Either or: -Slow, or -Fast")
		}
		if *slowF && *listF {
			die("Either or: -Slow, or -List")
		}
		if *listF && *fastF {
			die("Either or: -List, or -Fast")
		}
	*/

	if *dGoalMap {
		*dGoal = *dGoalMap
	}
	if *dFailMap {
		*dFail = *dFailMap
	}
	if *dCallMap {
		*dCall = *dCallMap
	}
	if *dLeafMap {
		*dLeaf = *dLeafMap
	}
	if *dleafMap {
		*dleaf = *dleafMap
	}

}

// ========================================================
