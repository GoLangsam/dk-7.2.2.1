// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s

// ========================================================

// CallBacks from Dance - as D.On.Next

// Spin is what keeps a dancer to turn to dance dancing or not
func (d *pace) Spin() {

	if d.On.Done != nil && d.On.Done() { return } // YES We have to return

	if d.On.Goal != nil && d.On.Goal() { // YES We have a solution
		if d.beatOnSpin { d.drums.Goal.Beat(d.level) } // ... we count our happiness
		return // and return
	}

	next, ok := d.On.Fail() // ... keep going?

	if !ok { // YES We have a failure/dead-end
		if d.beatOnSpin { d.drums.Fail.Beat(d.level) } // ... we count our suffering
		return // and return

	} else { // YES We have to go on goaling
		if d.beatOnSpin { d.drums.Call.Beat(d.level) } // ... we count our effort
	}

	d.level++
	d.On.Next(next)
	d.level--
}

// SpinBeatless keeps a dancer turning to dance dancing or not - but without beating
func (d *On) SpinBeatless() {

	if d.Done != nil && d.Done() { return } // YES We have to return
	if d.Goal != nil && d.Goal() { return } // YES We have a solution
	next, ok := d.Fail()                    // ... keep going?
	if !ok { return }                       // YES We have a failure/dead-end

	//	d.level++
	d.Next(next)
	//	d.level--
}

// SpinVeryFast keeps a dancer turning to dance dancing or not - but without beating, and without turning OnLeaf nor OnGoal
//  Note: this may be useful only for benchmarks / OnLeaf-counting
func (d *On) SpinVeryFast() {

	//	if d.Done() { return } // YES We have to return
	//	if d.Goal() { return } // YES We have a solution - but we don't tell anyone
	next, ok := d.Fail()           // ... keep going?
	if !ok { return }              // YES We have a failure/dead-end

	//	d.level++
	d.Next(next)
	//	d.level--
}
