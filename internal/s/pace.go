// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x"
)

// pace is what dancer loves to follow.
type pace struct {
	On

	drums
	level x.Index

	beatOnSpin bool
}

// ========================================================
