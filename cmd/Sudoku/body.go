// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/GoLangsam/do"

	"github.com/GoLangsam/dk-7.2.2.1/data"
	"github.com/GoLangsam/dk-7.2.2.1/internal/d" // dancing
	"github.com/GoLangsam/dk-7.2.2.1/internal/m" // problem matrix
)

func body(done do.Nok) {

	for _, hint := range []func() (string, data.SudokuHint){data.Hint28a, data.Hint28b, data.Hint29a, data.Hint29a67, data.Hint29b, data.Hint29c} {

		name, lines := data.Sudoku(hint())

		M := func() m.M { return m.Problem(name, lines...).Matrix() }

		for _, recur := range []bool{true, false} {
			if done.Do() {
				return
			}

			var a d.D
			if recur {
				a = d.RecursiveX(M)
			} else {
				a = d.AlgorithmX(M)
			}

			_ = a.Settings(
				(&a.On.Skip).Add(done.Do),
				getFlagOption(),
			)

			a.Search()
		}
	}
}
