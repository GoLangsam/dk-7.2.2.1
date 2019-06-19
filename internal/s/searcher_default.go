// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s

import (
	"fmt"

	"github.com/GoLangsam/dk-7.2.2.1/internal/x"
)

// ========================================================

// Default methods

func (a *Searcher) onDone() bool { // Do we have to return?
	if a.logOnDone { // TODO: What to print?
		// d.Stacker.Top().PrintValue("Stack-Top")
		// l.PrintValue()
		// l.PrintAways()
	}
	return false
}

// ========================================================

func (a *Searcher) onGoal() bool { // Do we have a solution?
	root := a.ItemS[0]
	if root.Next == 0 { // YES We have a solution
		if a.logOnGoal { // ... we may show it
			fmt.Println("Solution:")
			show := func(i x.Index) { fmt.Print(a.m.NameS[a.OptaS[i].Root], tab) }
			do := a.m.Do(show)
			for _, opta := range a.Stack {
				// show(opta)
				do.ForEachLineNext(opta)
				fmt.Println()
			}
		}
		return true
	}

	return false
}

// ========================================================

// chooseMRV searches the next item to be considered (if any).
// It implements the MRV heuristic.
func (a *Searcher) chooseMRV() (here x.Index, found bool) {

	Size := x.Index(len(a.OptaS)) // larger than anything we'll find.

	root := a.ItemS[0]
	for curr, size := root.Next, x.Index(0); curr != 0; curr = a.ItemS[curr].Next {

		size = a.OptaS[curr].Root
		// TODO: the "non-sharp/sharp preference"-Heuristics
		// if a.NameS[curr] does/doesn't start with `#` {
		//	size = size + len(a.Optas.MarkS) - 1 // #-of-options
		// }
		if size == 0 {
			Size, here, found = -1, -1, false
			break
		}
		if size < Size {
			Size, here, found = size, curr, true
		}
	}

	if a.logChoice { // ... we may show it
		fmt.Println("Chosen:", tab, here, tab, found, tab)
	}
	return
}

// ========================================================
