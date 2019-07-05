// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/do"
)

// ===========================================================================

// Settings applies the Setters
// and returns its undo function
// which, when evaluated, fully
// restores the previous setup.
func (a *D) Settings(doit ...do.Option) do.Opt {
	return do.Options(a, doit...)
}

// ===========================================================================

// WrapIt returns a do.It which wraps the given functions applied to a.
func (a *D) WrapIt(fs ...func(*D)) do.It {
	switch len(fs) {
	case 0:
		return nil
	case 1:
		return func() { fs[0](a) }
	default:
		return func() {
			for _, f := range fs {
				f(a)
			}
		}
	}
}

// ===========================================================================

// WrapOk returns a do.Ok which wraps the given functions applied to a.
func (a *D) WrapOk(fs ...func(*D) bool) do.Ok {
	switch len(fs) {
	case 0:
		return nil
	case 1:
		return func() bool { return fs[0](a) }
	default:
		return func() bool {
			for _, f := range fs {
				if !f(a) {
					return false
				}
			}
			return true
		}
	}
}

// ===========================================================================

// WrapNok returns a do.Ok which wraps the given functions applied to a.
func (a *D) WrapNok(fs ...func(*D) bool) do.Nok {
	switch len(fs) {
	case 0:
		return nil
	case 1:
		return func() bool { return fs[0](a) }
	default:
		return func() bool {
			for _, f := range fs {
				if f(a) {
					return true
				}
			}
			return false
		}
	}
}

// ===========================================================================
