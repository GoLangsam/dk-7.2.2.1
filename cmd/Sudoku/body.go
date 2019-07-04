// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/GoLangsam/do"

	"github.com/GoLangsam/dk-7.2.2.1/data"
	"github.com/GoLangsam/dk-7.2.2.1/internal/m" // problem matrix
	"github.com/GoLangsam/dk-7.2.2.1/internal/s" // searcher
)

func body(done do.Nok) {
	for _, d := range []struct {
		name string
		data [][]string
	}{
		{name: "SmallMatrix", data: data.SmallMatrix()},
		{name: "Sudoku 29a", data: data.Sudoku(data.Hint29a())},
		{name: "Sudoku 29b", data: data.Sudoku(data.Hint29b())},
		{name: "Sudoku 29c", data: data.Sudoku(data.Hint29c())},
	} {

		if done.Do() {
			break
		}

		M := func() m.M { return m.Problem(d.name, d.data...).Matrix() }

		for _, recur := range []bool{true, false} {
			if done.Do() {
				break
			}

			var a *s.S
			if recur {
				a = s.RecursiveX(M)
			} else {
				a = s.AlgorithmX(M)
			}

			_ = a.Settings(
				(&a.D.On.Skip).Add(done.Do),
				getFlagOption(),
			)

			a.Search()
		}
	}
}
