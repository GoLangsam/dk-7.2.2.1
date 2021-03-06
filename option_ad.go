// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dl

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/d" // dancing
)

// ===========================================================================

// aD returns any as a *d.D
// or panics, iff any has some other 'strange' type.
//
// Use of aD in every do.Option function
// allows to apply such on a (*)d.D.
func aD(any interface{}) *d.D {
	switch t := any.(type) {
	case *d.D:
		return t
	case d.D:
		return &t
	default:
		dif("aD: this option is for type %T - not for some *d.D", any)
	}
	return nil
}

// ===========================================================================
