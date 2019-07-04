// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

import (
	"github.com/GoLangsam/container/oneway/drum"
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
	"github.com/GoLangsam/do"
)

// ===========================================================================

type tach struct {
	x.ItemS    // the items
	x.OptaS    // the item-headers and the (spaced) options
	*drum.Drum // the update Drum - incremented in DoHide
	*do.It     // the update action
}

// TachX implements d.Tacher.
type tachX struct {
	tach
}

// TachC implements d.Tacher.
type tachC struct {
	tach
}

// ===========================================================================
