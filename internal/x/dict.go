// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x

import (
	"fmt"
)

// ===========================================================================

// Dict maps any Name to its Index.
type Dict map[Name]Index

// ===========================================================================

// LearnOnce registers some name and its index
// and panics iff such is already known.
func (a Dict) LearnOnce(name Name, index Index) Dict {
	if i, ok := a[name]; ok {
		die(fmt.Sprintf("Dict: name `%v` is duplicate - first seen at %v - now at %v!", name, i, index))
	}

	a[name] = index
	return a
}

// MustKnow returns the index for a given name
// and panics iff such is not yet known.
func (a Dict) MustKnow(name Name) Index {
	i, ok := a[name]
	if !ok {
		die(fmt.Sprintf("Dict: name `%v` is unknown!", name))
	}
	return i
}

// ===========================================================================
