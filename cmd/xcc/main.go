// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/GoLangsam/dk-7.2.2.1"
	"github.com/GoLangsam/dk-7.2.2.1/data"
)

const tab = "\t"

func main() {

	_ = data.SmallMatrix()

	m := data.SmallMatrix()

	if false {
		fmt.Println(len(m.NameS), tab, m.NameS)
		fmt.Println(len(m.ItemS), tab, m.ItemS)
		fmt.Println(len(m.OptaS), tab, m.OptaS)
	}

	if false {
		do := m.Do(func(i int) { fmt.Print(tab, m.NameS[i]) })

		do.ForEachNext(0) // 	a	b	c	d	e	f	g
		fmt.Println()     // .
		do.ForEachNext(3) // 	d	e	f	g		a	b
		fmt.Println()     // .
		do.ForEachPrev(0) // 	g	f	e	d	c	b	a
		fmt.Println()     // .
		do.ForEachPrev(3) // 	b	a		g	f	e	d
		fmt.Println()     // .
	}

	if false {
		do := m.Do(func(i int) { fmt.Print(m.NameS[m.OptaS[i].Root], tab) })

		do.ForEachOtherNext(len(m.ItemS) + 1) // e
		fmt.Println()                         // .
		do.ForEachOtherNext(len(m.ItemS) + 2) // c
		fmt.Println()                         // .
		//do.ForEachOtherNext(len(m.ItemS) + 3) // log.Panic
		//fmt.Println()                         // .
		do.ForEachOtherNext(len(m.ItemS) + 4) // d	g
		fmt.Println()                         // .
		do.ForEachOtherNext(len(m.ItemS) + 5) // g	a
		fmt.Println()                         // .
		do.ForEachOtherNext(len(m.ItemS) + 6) // a	d
		fmt.Println()                         // .
	}

	{
		m := data.FourQueens()
		fmt.Println(len(m.NameS), tab, m.NameS)
		fmt.Println(len(m.ItemS), tab, m.ItemS)
		fmt.Println(len(m.OptaS), tab, m.OptaS)

		dl.Search(m)
		// dl.Search(data.FourQueens())
		fmt.Println("Search finished!")
	}

}
