// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:pattern "github.com/GoLangsam/container/oneway/drum"

// drum.go provides a simple counter (with names based on musical methaphores)

package x

import (
	"fmt"
	"sort"
	//	"sync"
)

type counter map[Index]int64

// Drum is named counter
type Drum struct {
	Nam    string
	Cnt    int64
	Map    counter
	UseMap bool
	//	sync.Mutex
}

// NewDrum returns a new initialized drum.
func NewDrum(nam string, cap int) Drum {
	return Drum{
		Nam: nam,
		Map: make(map[Index]int64, cap),
	}
}

// Beat increments b.Cnt by one.
// And iff b.Verbose then b.Map[cur] get incremented as well.
func (b *Drum) Beat(cur Index) {
	//	b.Lock()
	//	defer b.Unlock()
	b.Cnt++
	if b.UseMap {
		b.Map[cur]++
	}
}

// Sort returns the keys of b.Map in a sorted slice
func (b *Drum) Sort() []int {
	//	b.Lock()
	//	defer b.Unlock()
	var keys sort.IntSlice
	for key := range b.Map {
		keys = append(keys, int(key))
	}
	keys.Sort() // Note: see also sort.Ints( []int )
	return keys
}

// Print prints the counter, if it's not empty, as lines of tab-terminated cells.
// And iff b.verbose then b.Map is printed in ascending order of its keys.
func (b *Drum) Print() {
	//	b.Lock()
	//	defer b.Unlock()
	if b.Cnt < 1 { // do not print empty counter
		return
	}
	fmt.Printf("%s\t% 9d\t\n", b.Nam, b.Cnt)
	if b.UseMap {
		for _, key := range b.Sort() {
			fmt.Printf("%6d\t% 9d\t\n", key, b.Map[Index(key)])
		}
	}
}
