package d

import (
	"github.com/GoLangsam/dk-7.2.2.1/internal/x" // all we need
)

// ===========================================================================

func (l *L) algorithmX(
	d tacher,
	on *On,
	root *x.Item,
	next chooser,
	mainS []x.Index,
) {

	var (
		i x.Index // main col - found by Choice, for DoCover & UnCover
		v x.Index // option cell - found on Stack.Top(), for DoCoverOthers & UnCoverOthers
	)

X2: // [Enter level] =============================================
	if on.Skip != nil && on.Skip.Do() { // => X8 - skip level
		goto X8 // >>>>>>>>>>
	}
	if root.Next == 0 { // => X8 (all items have been covered),
		if on.Goal != nil {
			on.Goal.Do()
		} // we have a Solution!
		goto X8 // >>>>>>>>>>
	}
	i = next(root) // X3. [Choose c]

	if l.optaS[i].Next == i { // => X7 - deadend
		if on.Fail != nil {
			on.Fail.Do()
		} // we have a Dead end!
		goto X7 // >>>>>>>>>>
	}

	{
		d.DoCover(i) // Inline ========================================
	}
	l.CellS[l.Index] = l.optaS[i].Next // delayed a.Stack.Push()

X5: // [Try x[l].] ===========================================================
	v = l.CellS[l.Index] // a.Stack.Pop()
	if v == i {          // => X7 - no more
		goto X7 // (we've tried all options for i).
	} // >>>>>>>>>>
	d.DoCoverOthers(v) // Inline ========================================
	l.Index++          // Incr Level
	goto X2            // >>>>>>>>>>
X6: // [Try again] ===========================================================
	v = l.CellS[l.Index] // cell.Stack.Top()

	d.UnCoverOthers(v) // Inline ========================================

	i = l.optaS[v].Root
	l.CellS[l.Index] = l.optaS[v].Next
	goto X5 // >>>>>>>>>>
X7: // [Backtrack] ===========================================================
	{
		d.UnCover(i) // Inline ========================================
	}

X8: // [Leave level] =========================================================
	if l.Index == 0 { // => X9
		goto X9 // >>>>>>>>>>
	}
	l.Index-- // Decr Level
	goto X6   // >>>>>>>>>>

X9: // we are done
	return
}
