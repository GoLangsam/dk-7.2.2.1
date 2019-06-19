// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:pattern "github.con/GoLangsam/container/oneway/list/form"

package s

// doFn is a self referential function.
//
// returning it's own undo function.
type doFn func() doFn

// Setter (as returned by many functions)
// is a function which
// modifes some *Searcher when applied to it
// and returns its own undo function (as doFn),
//  which, when applied, returns it's redo function (as doFn),
//  which, when applied, returns it's redo undo function (as doFn),
//  which, when applied, returns it's redo undo undo function (as doFn),
//  which, when applied, returns ...
//  (and so on - as doFn is a self referential function).
//
// Hint: Apply multiple Setters at once, using function (or method) Form(...).
//
// Note: Inspired by:
//   - http://commandcenter.blogspot.com.au/2014/01/self-referential-functions-and-design.html
//   - https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
//   - https://www.calhoun.io/using-functional-options-instead-of-method-chaining-in-go/
//
// (Undo is only supported for the last Setter passed in these samples.)
//
// This implementation of Settings(...) provides full undo.
type Setter func(*Searcher) doFn

// Settings applies the Setters
// and returns its undo function
// which, when evaluated, fully
// restores the previous state.
func (a *Searcher) Settings(doit ...Setter) doFn {

	prev := make([]doFn, 0, len(doit))

	for i := range doit {
		prev = append(prev, doit[i](a))
	}

	return func() doFn {
		return undo(prev...)
	}
}

// undo applies the given doit functions in reverse order
// and returns it's own undo.
func undo(doit ...doFn) doFn {

	switch len(doit) {
	case 1:
		return doit[0]()
	case 0:
		return func() doFn {
			return undo()
		}
	default:
		prev := make([]doFn, 0, len(doit))
		for i := len(doit) - 1; i >= 0; i-- {
			prev = append(prev, doit[i]())
		}
		return func() doFn {
			return undo(prev...)
		}
	}
}
