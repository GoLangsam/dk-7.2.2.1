// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:pattern "github.com/GoLangsam/container/oneway/stack/stack.go"

/*
Package stack provides a normal (=non-concurrency-safe) stack
for Index

Note: the very crisp implementation was found in cmd/go/pkg.go importStack
*/

package x

// Stack implements a normal (=non-concurrency-safe) stack
// based on a slice of Index, and never shrinking
type Stack []Index

// NewStack returns a new empty stack with given initial capacity
func NewStack(cap int) Stack {
	return make([]Index, 0, cap)
}

// Push sth onto the current stack
func (s *Stack) Push(a Index) {
	//	s.Lock()
	//	defer s.Unlock()

	*s = append(*s, a)
}

// Pop sth off the current stack
func (s *Stack) Pop() Index {
	//	s.Lock()
	//	defer s.Unlock()

	p := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return p
}

// Drop pops sth silently off the current stack
func (s *Stack) Drop() {
	//	s.Lock()
	//	defer s.Unlock()

	*s = (*s)[0 : len(*s)-1]
}

// Top returns the top of the current stack
func (s *Stack) Top() Index {
	//	s.Lock()
	//	defer s.Unlock()

	return (*s)[len(*s)-1]
}

// Get returns a copy (clone) of the current stack
func (s *Stack) Get() []Index {
	//	s.Lock()
	//	defer s.Unlock()

	//	return append([]Index{}, *s...)
	var stack = make([]Index, len(*s))
	copy(stack, *s)
	return stack
}

// Len returns the length of the current stack
func (s *Stack) Len() int {
	//	s.Lock()
	//	defer s.Unlock()

	return len(*s)
}
