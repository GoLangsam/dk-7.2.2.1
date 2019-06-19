package d

func (a *D) AlgX() {

	var c Index // main col - found by Choice, for DoCover & UnCover
	var v Index // option cell - found on Stack.Top(), for DoCoverOthers & UnCoverOthers

	a.Stack = make([]Index, len(a.ItemS))

X2: // [Enter level l = a.Index]

	// TODO: report iff Stack is too short; panic (easy), or extend it (not so easy)!

	// TODO: report OnCall - Good time to call home and check for abort and setters

	// TODO: report OnLeaf: No need here !!! :-)

	if a.OptaS[0].Root == 0 { // (hence all items have been covered),
		// TODO: report OnGoal: visit the solution that is specified by x[0], x[1], ..., x[l-1] and
		goto X8
	}

	// X3. [Choose c]
	c = a.On.Here()

	if a.OptaS[c].Next == c {
		// TODO: report OnFail: we have a Dead end!
		goto X7
	}

	// X4. [Cover c]
	{
		a.DoCover(c) // Inline ========================================
	}
	a.Stack[a.Index] = a.OptaS[c].Next // delayed a.Stack.Push()

X5: // [Try x[l].]
	v = a.Stack[a.Index] // a.Stack.Pop()
	if v.Root == c {
		goto X7 // (we've tried all options for i).
	} // Otherwise set
	a.DoCoverOthers(v) // Inline ========================================

	a.Index++
	goto X2

X6: // [Try again.]
	v = a.Stack[a.Index]
	a.UnCoverOthers(v) // Inline ========================================

	c = a.OptaS[v].Root
	a.Stack[a.Index] = a.OptaS[v].Next
	goto X5

X7: // [Backtrack.]
	{
		a.UnCover(c) // using (14).
	}

X8: // [Leave level l].
	if a.Index == 0 {
		return
	} // Otherwise
	a.Index--
	goto X6

}
