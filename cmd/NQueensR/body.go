// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"strconv"

	"github.com/GoLangsam/do"

	"github.com/GoLangsam/dk-7.2.2.1/data"
	"github.com/GoLangsam/dk-7.2.2.1/internal/m" // problem matrix
	"github.com/GoLangsam/dk-7.2.2.1/internal/s" // searcher
)

func body(done do.Nok) {
	for N := *beg; N <= *end && !done.Do(); N++ {

		name := strconv.Itoa(N) + "-QueensR"
		for _, recur := range []bool{true, false} {
			if done.Do() {
				break
			}

			M := func() m.M { return m.Problem(name, data.NQueensR(N)...).Matrix() }
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
