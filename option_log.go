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
