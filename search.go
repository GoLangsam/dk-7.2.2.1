// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package dl allows to solve the exact cover problem via dancing links.
//
// The implementation is spread across internal packages:
//
//  x provides basic constituents (Names, Items, Optas, Stack, Drum ...).
//
//  m implements M, the matrix representing a problem.
//
//  d provides D, the dancer, dancing on the links - pretty fast
//    using components of M vital for its graceful dance.
//    This is the core of the algorithm - the heartbeat!
//
//  p allows to build a matrix M given lists of symbols.
//
//  s provides a Searcher who -given some M- creates a suitable D
//    and manages its heartbeats pace, operation and output.
//    Its many setters permit finetuning on the fly.
//    Client-provided functions can be plugged in.
package dl

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/m" // problem matrix
	"github.com/GoLangsam/dk-7.2.2.1/internal/p" // problem matrix builder
	"github.com/GoLangsam/dk-7.2.2.1/internal/s" // searcher
)

// Matrix is (a pointer to) a problem matrix M
type Matrix = *m.M

// Problem wraps p.Problem - see there for detailed info.
func Problem(lines ...[]string) *p.P {
	return p.Problem(lines...)
}

// Items wraps p.Items - see there for detailed info.
func Items(items ...string) *p.P {
	return p.Items(items...)
}

// Search is an example.
// It shows how to obtain a s.Searcher given a problem Matrix
// and to have it search and print.
func Search(M Matrix) {
	var v, vg, rh, vd, vb, vc bool // as in tees/cmd/NQueensR
	_, _ = vb, rh                  // d.Dancer = dancing.New(vb, rh, false)
	// pace.VerboseTurn(vt) // vt not defined,
	var vs bool // TODO: support in cmd/NQueensR
	// vc = true
	vd = true
	vg = true
	vs = true

	var vdd bool // new: for dance drum
	// vdd = true

	var vdg, vdf, vdc, vdl bool // new: for each drum
	// vdg, vdf, vdc, vdl = vd, vd, vd, vd

	s.New(M,
		s.LogOnDone(v),
		s.LogOnGoal(vg),
		s.LogChoice(vc),
		s.VerboseOnSpin(vs),
		s.VerboseDrums(vd),
		s.VerboseDrumsGoal(vdg),
		s.VerboseDrumsFail(vdf),
		s.VerboseDrumsCall(vdc),
		s.VerboseDrumsLeaf(vdl),
		s.CountUpdatesPerCell(vdd),
		// ...
	).Search().Print()
}
