// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/GoLangsam/dk-7.2.2.1"
	"github.com/GoLangsam/dk-7.2.2.1/data"
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

const tab = "\t"

func main() {

	m := dl.Problem("SmallMatrix", data.SmallMatrix()...)

	if false {
		fmt.Println(len(m.NameS), tab, m.NameS)
		fmt.Println(len(m.ItemS), tab, m.ItemS)
		fmt.Println(len(m.OptaS), tab, m.OptaS)
	}

	if false {
		do := m.Do(func(i x.Index) { fmt.Print(tab, m.NameS[i]) })

		do.ForEachNext(x.Index(0)) // 	a	b	c	d	e	f	g
		fmt.Println()              // .
		do.ForEachNext(x.Index(3)) // 	d	e	f	g		a	b
		fmt.Println()              // .
		do.ForEachPrev(x.Index(0)) // 	g	f	e	d	c	b	a
		fmt.Println()              // .
		do.ForEachPrev(x.Index(3)) // 	b	a		g	f	e	d
		fmt.Println()              // .
	}

	if false {
		do := m.Do(func(i x.Index) { fmt.Print(m.NameS[m.OptaS[i].Root], tab) })

		do.ForEachOtherNext(x.Index(len(m.ItemS) + 1)) // e
		fmt.Println()                                  // .
		do.ForEachOtherNext(x.Index(len(m.ItemS) + 2)) // c
		fmt.Println()                                  // .
		//do.ForEachOtherNext(x.Index(len(m.ItemS) + 3)) // log.Panic
		//fmt.Println()                         // .
		do.ForEachOtherNext(x.Index(len(m.ItemS) + 4)) // d	g
		fmt.Println()                                  // .
		do.ForEachOtherNext(x.Index(len(m.ItemS) + 5)) // g	a
		fmt.Println()                                  // .
		do.ForEachOtherNext(x.Index(len(m.ItemS)) + 6) // a	d
		fmt.Println()                                  // .
	}

	if false {
		N := 80 // TODO: N = 80 is the max! Beyond => index overflow
		name := strconv.Itoa(N) + "-Queens"
		M := func() dl.Matrix { return dl.Problem(name, data.NQueens(N)...) }
		m := M()
		fmt.Println(len(m.NameS), tab, m.NameS)
		fmt.Println("Size:", tab, reflect.TypeOf(m).Size(), unsafe.Sizeof(m))
		s := int(unsafe.Sizeof(m.ItemS[0]))
		fmt.Println("Size ItemS:", tab, reflect.TypeOf(m.ItemS).Size(), len(m.ItemS), len(m.ItemS)*s)
		s = int(unsafe.Sizeof(m.OptaS[0]))
		fmt.Println("Size OptaS:", tab, reflect.TypeOf(m.OptaS).Size(), len(m.OptaS), len(m.OptaS)*s)
	}

	if true {
		for _, d := range []struct {
			name string
			data [][]string
		}{
			{name: "SmallMatrix", data: data.SmallMatrix()},
			{name: "Sudoku 29a", data: data.Sudoku(data.Hint29a())},
			//{name: "Sudoku 29b", data: data.Sudoku(data.Hint29b())},
			//{name: "Sudoku 29c", data: data.Sudoku(data.Hint29c())},
		} {

			M := func() dl.Matrix { return dl.Problem(d.name, d.data...) }
			for _, recur := range []bool{true, false} {
				dl.Search(M, recur)
}
			fmt.Println()
		}
	}

	if false {
		for N := 1; N < 14; N++ {

			name := strconv.Itoa(N) + "-Queens"
			for _, recur := range []bool{true, false} {
				M := func() dl.Matrix { return dl.Problem(name, data.NQueens(N)...) }
				dl.Search(M, recur)
			}
			fmt.Println()
		}
	}

}
