// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x

// ===========================================================================

// DoIf wraps (a pointer to) a Do func(int) and a boolean DoIt switch
// in order to facilitate conditional invocation via its DoIf(int) method..
//
// Intended use is for conditional logging, counting etc.
//
// The null value is useful: it's a nop.
type DoIf struct {
	Do
	DoIt bool
}

func (a *DoIf) DoIf(i int) {
	if a != nil && a.DoIt && a.Do.Do != nil {
		a.Do.Do(i)
	}
}

// ===========================================================================
