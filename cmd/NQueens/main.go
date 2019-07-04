// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/GoLangsam/do"
	"github.com/GoLangsam/do/cli/cancel"
)

func main() {
	flagParse()

	ctx, _ := cancel.WithCancel()

	var done do.Nok
	done = func() bool {
		select {
		case <-ctx.Done():
			return true
		default:
		}
		return false
	}

	body(done)
	if ctx.Err() != nil {
		fmt.Println("Early graceful exit - because:", ctx.Err().Error())
	}

	fmt.Println()

}
