// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "log"

// se_ is the logging function
var se_ = log.Print

// see is the logging function
var see = log.Println

// sef is the logging function
var sef = log.Printf

// di_ logs the fatal error
var di_ = log.Fatal

// die logs the fatal error
var die = log.Fatalln

// die logs the fatal error
var dif = log.Fatalf

var verbose bool

// qq_ logs iff verbose
func qq_(args ...interface{}) {
	if verbose {
		se_(args...)
	}
}

// qqq logs iff verbose
func qqq(args ...interface{}) {
	if verbose {
		see(args...)
	}
}

// qqf logs iff verbose
func qqf(format string, args ...interface{}) {
	if verbose {
		sef(format, args...)
	}
}

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
}
