// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

// ===========================================================================

// On consolidates functions relevant for controlling behaviour.
type On struct {
	Here func() int // mandatory - chooses (next) item to cover - use MRV heuristic per default
	Next func()     // mandatory for recursive version: Callback to controller
	Leaf func(int)  // optional: Called per Update with Level
}

// ===========================================================================
