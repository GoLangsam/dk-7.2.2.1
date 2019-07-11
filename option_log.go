// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dl

import (
	"fmt"
	"runtime"
	"time"

	"github.com/GoLangsam/do"
	"github.com/GoLangsam/do/ami"
)

// ===========================================================================

// LogSize logs a starting-line upon Init
// with the size of the problem representation.
func LogSize() do.Option {
	return func(any interface{}) do.Opt {
		a := aD(any)

		lI := uint(len(a.ItemS))
		lO := uint(len(a.OptaS))
		sI := uint(ami.TypeSize(a.ItemS[0]))
		sO := uint(ami.TypeSize(a.OptaS[0]))
		items := fmt.Sprint("Items:", tab, lI, "*", sI, "=", lI*sI)
		optas := fmt.Sprint("Optas:", tab, lO, "*", sO, "=", lO*sO)

		init := func() {
			see(a.Name, a.M.Name,
				tab, "starting",
				tab, items,
				tab, optas,
			)
		}
		undoInit := (&a.On.Init).Add(init)(a)

		return func() do.Opt {
			return undoInit
		}
	}
}

// ===========================================================================

// LogTime logs a finished-line upon Done
// with the time elapsed since On.Init.
//
// Note: runtime.GC is forced before the timer starts.
func LogTime() do.Option {
	return func(any interface{}) do.Opt {
		a := aD(any)

		var t time.Time
		init := func() { t = time.Now() }
		show := func() string { return fmt.Sprint(time.Since(t)) }
		done := func() {
			see(a.Name, a.M.Name,
				tab, "finished",
				tab, "Time:",
				tab, show())
		}

		undoInit := (&a.On.Init).Add(runtime.GC, init)(a)
		undoDone := (&a.On.Done).Add(done)(a)

		return func() do.Opt {
			return do.OptJoin(undoInit, undoDone)
		}
	}
}

// ===========================================================================

// LogProgress sends a progress line to the current logger
// upon everyNmillion (minimum = 10) updates
// which shows the completition (in permille)
// as well as the number of Grooves, DeadEnd and millions-of-Updates.
func LogProgress(everyNmillion int64) do.Option {
	return func(any interface{}) do.Opt {
		a := aD(any)

		if everyNmillion < 1 {
			everyNmillion = 1
		}
		million := int64(1000000)
		step := everyNmillion * million
		N := step - 1

		var goals, fails, leafs int64
		goal := func() { goals++ }
		fail := func() { fails++ }
		leaf := func() {
			leafs++
			if leafs > N {
				see(a.Permille(), "%o",
					tab, "Grooves:", goals,
					tab, "DeadEnd:", fails,
					tab, "Updates:", leafs/million, "Mio.",
				)
				N += step
			}
		}

		undoGoal := (&a.On.Goal).Add(goal)(a)
		undoFail := (&a.On.Fail).Add(fail)(a)
		undoLeaf := (&a.On.Leaf).Add(leaf)(a)
		return func() do.Opt {
			return do.OptJoin(undoGoal, undoFail, undoLeaf)
		}
	}
}
