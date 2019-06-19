// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package data

//go:pattern "github.com/GoLangsam/tees/list/test/IDs.go"

import (
	"fmt"
)

func width(anz int) int {
	// too lazy to use log ;-)

	switch {
	case anz > 999999999:
		return 10
	case anz > 99999999:
		return 9
	case anz > 9999999:
		return 8
	case anz > 999999:
		return 7
	case anz > 99999:
		return 6
	case anz > 9999:
		return 5
	case anz > 999:
		return 4
	case anz > 99:
		return 3
	case anz > 9:
		return 2
	default:
		return 1
	}
}

func getFormatWidthPaddingZeros(anz int) string {
	return "%0" + fmt.Sprintf("%d", width(anz)) + "d"
}

func getFormatWidthPaddingSpaces(anz int) string {
	return "% " + fmt.Sprintf("%d", width(anz)) + "d"
}

func getFormatWidth(prefix string, anz int) string {
	if prefix == "" {
		return "%s" + getFormatWidthPaddingSpaces(int(anz))
	}
	return "%s" + getFormatWidthPaddingZeros(int(anz))
}

// IDs returns a slice of size-adjusted IDs.
func IDs(prefix string, anz int) []string {

	var s = make([]string, 0, anz)
	var f = getFormatWidth(prefix, anz)

	for i := 0; i < anz; i++ {
		id := fmt.Sprintf(f, prefix, i+1)
		s = append(s, id)
	}

	return s
}
