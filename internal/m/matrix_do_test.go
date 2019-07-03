// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package m_test

import (
	"fmt"

	"github.com/GoLangsam/dk-7.2.2.1/data"
	"github.com/GoLangsam/dk-7.2.2.1/internal/x"
)

// ===========================================================================

func Example_forEachNext() {
	a := data.SmallMatrixM()

	showName := func(i x.Index) { fmt.Print(" ", a.NameS[int(i)]) }

	for _, root := range a.Items.MarkS {
		a.Do(showName).ForEachNext(root)
		fmt.Println()
	}

	// Output:
	//  a b c d e f g
}

func Example_forEachPrev() {
	a := data.SmallMatrixM()

	showName := func(i x.Index) { fmt.Print(" ", a.NameS[int(i)]) }

	for _, root := range a.Items.MarkS {
		a.Do(showName).ForEachPrev(root)
		fmt.Println()
	}

	// Output:
	//  g f e d c b a
}

// ===========================================================================

func Example_forEachLineNext() {
	a := data.SmallMatrixM()

	showName := func(i x.Index) { fmt.Print(" ", a.NameS[a.OptaS[int(i)].Root]) }

	for i, root := range a.Optas.MarkS {
		fmt.Print("#", i, ":") //, a.OptaS[root].Prev, a.OptaS[root].Next)
		a.Do(showName).ForEachLineNext(root)
		fmt.Println()
	}

	// Output:
	// #0:
	// #1: c e
	// #2: a d g
	// #3: b c f
	// #4: a d f
	// #5: b g
	// #6: d e g
}

func Example_forEachLinePrev() {
	a := data.SmallMatrixM()

	showName := func(i x.Index) { fmt.Print(" ", a.NameS[a.OptaS[int(i)].Root]) }

	for i, root := range a.Optas.MarkS {
		fmt.Print("#", i, ":") //, a.OptaS[root].Prev, a.OptaS[root].Next)
		a.Do(showName).ForEachLinePrev(root)
		fmt.Println()
	}

	// Output:
	// #0: e c
	// #1: g d a
	// #2: f c b
	// #3: f d a
	// #4: g b
	// #5: g e d
	// #6:
}

/*

func Example_forEachOptaPrev() {
	a := data.SmallMatrixM()

	for _, root := range a.Items.MarkS {
		if a.OptaS[root].Prev != 0 {
			a.Do(showName(a)).ForEachOptaPrev(root)
			fmt.Println()
		}
	}

	// Output:
	// g	f	e	d	c	b	a
}

/*

func Example_forEachLineNext() {
	a := data.SmallMatrix()

	for _, root := range a.Optas.MarkS {
		a.Do(showOptaName(a)).ForEachLineNext(root)
		fmt.Println()
	}

	// Output:
	// g	f	e	d	c	b	a
}

*/
