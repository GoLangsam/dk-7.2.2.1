// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package m

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x"
)

// ===========================================================================

// Do returns a Do (which wraps matrix data and the given do-function and)
// which provides iterating methods ForEach.... or some conditional DoIf.
//
// Note: Do for a M wraps an empty stack.
func (a *M) Do(f func(int)) x.Do {
	return x.Do{&a.Items, &a.Optas, &x.Stack{}, f}
}

// ===========================================================================

func (a *M) DoIf(f func(int)) *x.DoIf {
	return &x.DoIf{Do: x.Do{Do: f}}
}

// ===========================================================================
