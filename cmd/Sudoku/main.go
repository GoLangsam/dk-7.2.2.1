// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/GoLangsam/do/cli/cancel"
)

func main() {
	flagParse()

	ctx, _ := cancel.WithCancel()
	doneFn := cancel.Done(ctx)

	body(doneFn)

	if ctx.Err() != nil {
		see("Early graceful exit!", tab, "reason:", tab, ctx.Err().Error())
	}

	see()

}
