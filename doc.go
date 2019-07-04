// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package dl allows to solve any exact cover problem via dancing links.
//
// The implementation is spread across internal packages:
//
//  x provides basic constituents (Names, Items, Optas, Dict ...).
//
//  m implements M, the matrix representing a named problem
//    built from a given list of lists of symbols.
//
//  d provides D, a dancer dancing its graceful dance -pretty fast-
//    on the links of the given M which implements Searcher.
//    This is the core of the algorithm - the heartbeat!
//
//  s provides a Searcher who invokes a suitable D,
//    manages its heartbeat, operation and output,
//    given a way to build some M.
//    Many options permit finetuning on the fly:
//    plug-in functions as needed.
package dl
