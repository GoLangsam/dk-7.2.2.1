// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x

import (
	"fmt"
	"log"
)

// ===========================================================================

// Dict
type Dict map[string]Index

// ===========================================================================

// LearnOnce registers some name and its index
// and panics iff such is already known.
func (a Dict) LearnOnce(name string, index Index) Dict {
	if i, ok := a[name]; ok {
		log.Panic(fmt.Sprintf("Dict: name `%v` is duplicate - first seen at %v - now at %v!", name, i, index))
	}

	a[name] = index
	return a
}

// MustKnow returns the index for a given name
// and panics iff such is not yet known.
func (a Dict) MustKnow(name string) Index {
	i, ok := a[name]
	if !ok {
		log.Panic(fmt.Sprintf("Dict: name `%v` is unknown!", name))
	}
	return i
}

// ===========================================================================
