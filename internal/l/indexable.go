// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package l

// ===========================================================================

// Indexable has to be implemented by any collection of Node
// (which is typically some slice-type).
//
// The intention is to allow access to individual elements.
type Indexable interface {
	I(int) Node
}

// ===========================================================================
