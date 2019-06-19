// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x"
)

// ===========================================================================

// On consolidates functions relevant for controlling behaviour.
type On struct {
	Done func() bool            // YES We have to return
	Goal func() bool            // YES We have a solution
	Fail func() (x.Index, bool) // YES We have to go on goaling
	Next func(x.Index)          // and We have to dance here now
}

// ===========================================================================
