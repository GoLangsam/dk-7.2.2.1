// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dl

import (
	"fmt"

	"github.com/GoLangsam/do"
	// "github.com/GoLangsam/dk-7.2.2.1/internal/d"
	"github.com/GoLangsam/dk-7.2.2.1/internal/m" // problem matrix
	"github.com/GoLangsam/dk-7.2.2.1/internal/s" // searcher
)

// Matrix is a problem matrix M
type Matrix = m.M

// Problem wraps m.New - see there for detailed info.
func Problem(name string, lines ...[]string) Matrix {
	return *m.Problem(name, lines...).Matrix()
}

// as in tees/cmd/NQueensR
func getOption(v, vg, rh, vd, vb, vc, vdd, vdg, vdf, vdc, vdl bool) do.Option {
	// d.Dancer = dancing.New(vb, rh, false)
	// new: vdg, vdf, vdc, vdl for each drum
	// new: vdd for dance drum

	options := []do.Option{}

	if vg {
		options = append(options,
			// OnGoal(d.ShowGoal),
		)
	}

	if v {
		options = append(options,
			ShowSize(),
			ShowTime(),
		)
	}

	if vd {
		options = append(options,
			CountGrooves(vdg), // Count solutions
			CountDeadEnd(vdf), // Count dead-ends
			CountNiveaus(vdc), // Count effort
			CountUpDates(vdl), // Count updates
			CountUpdates(vdd), // Count updates per cell
			//s.LogChoice(vc),
			//s.VerboseOnSpin(vs),
			// ...
			/*
				,
			*/
		)
	}

	return do.OptionJoin(options...)
}

// Search is an example.
// It shows how to obtain a s.Searcher given a problem Matrix
// and to have it search and print.
func Search(M func() Matrix, recur bool) {
	var (
		v, vg, rh, vd, vb, vc bool // as in tees/cmd/NQueensR
		vdd                   bool // new: vdd for dance drum
		vdg, vdf, vdc, vdl    bool // new: vdg, vdf, vdc, vdl for each drum
	)
	_, _, _, _, _, _ = v, vg, rh, vd, vb, vc
	v = true
	// vc = true
	vd = true
	vg = true // show Goals
	//vdg, vdf, vdc, vdl, vdd = vd, vd, vd, vd, vd // useMap
	vdl, vdd = false, false // no updates map

	var a *s.S
	if recur {
		a = s.RecursiveX(M)
	} else {
		a = s.AlgorithmX(M)
	}

	_ = a.Settings(
		(&a.D.On.Init).Add(func() { fmt.Println(a.Name, tab, "starting", tab, "recur =", recur) }),
		getOption(v, vg, rh, vd, vb, vc, vdd, vdg, vdf, vdc, vdl),
	)

	a.Search()

}
