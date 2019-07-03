// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s

import (
	"github.com/GoLangsam/do"
	"github.com/GoLangsam/dk-7.2.2.1/internal/d" // dancing
	"github.com/GoLangsam/dk-7.2.2.1/internal/m" // problem matrix
)

// ===========================================================================

// S consolidates what's needed to Search.
type S struct {
	*d.D // the Dancer: can Dance - implements Searcher
}

// ===========================================================================

// Settings applies the Setters
// and returns its undo function
// which, when evaluated, fully
// restores the previous state.
func (a *S) Settings(doit ...do.Option) do.Opt {
	return do.Options(a, doit...)
}

// ===========================================================================

// RecursivX returns (a pointer to) a fresh Searcher
// with silent default settings.
func RecursiveX(M func() m.M, setters ...do.Option) *S {
	a := S{d.RecursiveX(M, false, true)}
	a.Settings(setters...) // user settings

	return &a
}

// RecursivC returns (a pointer to) a fresh Searcher
// with silent default settings.
func RecursiveC(M func() m.M, setters ...do.Option) *S {
	a := S{d.RecursiveX(M, true, true)}
	a.Settings(setters...) // user settings

	return &a
}

// AlgorithmX returns (a pointer to) a fresh Searcher
// with silent default settings.
func AlgorithmX(M func() m.M, setters ...do.Option) *S {
	a := S{d.AlgorithmX(M, false, true)}
	a.Settings(setters...) // user settings

	return &a
}

// AlgorithmC returns (a pointer to) a fresh Searcher
// with silent default settings.
func AlgorithmC(M func() m.M, setters ...do.Option) *S {
	a := S{d.AlgorithmX(M, true, true)}
	a.Settings(setters...) // user settings

	return &a
}

// ===========================================================================
