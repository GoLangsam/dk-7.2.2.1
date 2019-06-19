// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package solver

import (
	"container/list"
	"fmt"

	"github.com/GoLangsam/dk-7.2.2.1/internal/m" // problem matrix
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all You need
)

type Summary struct {
	// all kinds of drums
}

type Solution struct {
	*m.M
	x.Stack
	open // # of columns still open
}

// ===========================================================================
// Beg of Request

// Request is a function to be applied and channel on which to return the result.
type Request struct {
	fn func() Summary // operation to perform
	c  chan Summary   // channel on which to return result
}

// Beg of Fake
func workFn() (a Summary)      { return }
func furtherProcess(a Summary) {}

func requester(work chan<- Request) {
	cha := make(chan Summary)
	for {
		// time.Sleep ....
		work <- Request{workFn, cha} // send a work request
		result := <-cha              // wait for answer
		furtherProcess(result)
	}
}

func process() {
	requester(New(10))
}

// End of Fake

// End of Request
// ===========================================================================

// ===========================================================================
// Beg of Worker

type Worker struct {
	requests chan Request // work to do (a buffered channel)
}

// work is called from New()!
func (w *Worker) work(done chan<- *Worker) {
	for {
		req := <-w.requests // get requests from load balancer
		req.c <- req.fn()   // do the work and send the answer back to the requestor
		done <- w           // tell load balancer a task has been completed by worker w.
	}
}

// End of Worker
// ===========================================================================

// ===========================================================================
// Beg of Pool

// Pool is a list of (pointers to) Worker
type Pool struct {
	*list.List
}

// Len reports the number of elements in the heap.
func (p *Pool) Len() int { return (p.List.Len()) }

// Push add v as element @ [ Len() ].
func (p *Pool) Push(w *Worker) {
	p.List.PushBack(w)
}

// Pop removes and returns the element @ [ Len() - 1 ].
func (p *Pool) Pop() (v interface{}) {
	w := v.(*Worker)
	*p, v = (*p)[:p.Len()-1], (*p)[p.Len()-1]
	return
}

// End of Pool
// ===========================================================================

// ===========================================================================
// Beg of Balancer

// Balancer has a Pool of Workers and a channel for Workers having finished
type Balancer struct {
	pool Pool
	done chan *Worker
}

// New returns a receive-only request channel
// processed by `cap` balanced workers
func New(cap int) chan<- Request {
	b := &Balancer{
		pool: list.New()
		done: make(chan *Worker),
	}

	for i := 0; i < cap; i++ { // populate the worker pool
		work := make(chan Request) // work to receive
		w := Worker{work}          // by worker w
		b.pool.PushBack(w)         // as pool[i]
		go w.work(b.done)          // launch worker to work
	}

	work := make(chan Request)
	go b.balance(work)

	return work
}

// balance the work
func (b *Balancer) balance(work <-chan Request) {
	
	for b.pool.Len() > 0 {

		for e := b.pool.Front(); e != nil, e = b.pool.Next() { // worker

			select {
			case req := <-work: // request received
				 // forward request to a worker
				w := e.Value.(*Worker)
				w.requests <- req                // ... is assigned the task
				b.pool.List.PushBack(w)            // put it back in the heap

			case w := <-b.done: // worker finished with a request
				b.completed(w)
			}
			b.print()
		}
	}
}

// completed is the Job: update the Heap
func (b *Balancer) completed(w *Worker) {
	// TODO: what to do?
}

func (b *Balancer) print() {
	total_pending := 0
	sumsq_pending := 0
	for e := b.pool.Front(); e != nil, e = b.pool.Next() { // worker
		pending := 1 // TODO: get a better number
		fmt.Printf("%d  ", pending)
		total_pending += pending
		sumsq_pending += pending * pending
	}
	fmt.Printf("| %d  ", total_pending)
	avg := float64(total_pending) / float64(b.pool.Len())
	variance := float64(sumsq)/float64(len(b.pool)) - avg*avg
	fmt.Printf("| %.2f %.2f\n", avg, variance)

}

// End of Balancer
// ===========================================================================
