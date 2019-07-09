// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package data

// SmallMatrix returns the small problem
func SmallMatrix() (string, [][]string) {

	return "SmallMatrix",
		[][]string{
			{"a", "b", "c", "d", "e", "f", "g"},
			{},
			{"c", "e"},
			{"a", "d", "g"},
			{"b", "c", "f"},
			{"a", "d", "f"},
			{"b", "g"},
			{"d", "e", "g"},
		}
}
