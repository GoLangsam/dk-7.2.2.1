// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package data

import (
	"github.com/GoLangsam/dk-7.2.2.1"
)

func SmallMatrix() dl.Matrix {

	return dl.Problem([][]string{
		[]string{"a", "b", "c", "d", "e", "f", "g"},
		[]string{},
		[]string{"c", "e"},
		[]string{"a", "d", "g"},
		[]string{"b", "c", "f"},
		[]string{"a", "d", "f"},
		[]string{"b", "g"},
		[]string{"d", "e", "g"},
	}...).Matrix()
}
