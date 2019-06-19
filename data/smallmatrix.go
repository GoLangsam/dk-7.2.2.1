// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package data

import (
	"github.com/GoLangsam/dk-7.2.2.1"
)

func SmallMatrix() dl.Matrix {

	return dl.Problem([][]string{
		{"a", "b", "c", "d", "e", "f", "g"},
		{},
		{"c", "e"},
		{"a", "d", "g"},
		{"b", "c", "f"},
		{"a", "d", "f"},
		{"b", "g"},
		{"d", "e", "g"},
	}...).Matrix()
}
