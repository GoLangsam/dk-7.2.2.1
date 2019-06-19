// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x"
)

// ===========================================================================

// Do returns a Do (which wraps matrix data and the given do-function and)
// which provides iterating methods ForEach.... or some conditional Doer.
//
// Note: Do for a Searcher returns the Do of the embedded D.
func (a *Searcher) Do(f func(x.Index)) x.Do {
	return a.D.Do(f)
}

// ===========================================================================

func (a *Searcher) DoIf(f func(x.Index)) *x.DoIf {
	return &x.DoIf{Do: x.Do{Do: f}}
}

// ===========================================================================
