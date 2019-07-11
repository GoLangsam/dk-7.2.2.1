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
) {

	var (
		i  x.Main  // main col - found by Choice, for DoCover & UnCover
		v  x.Index // position of option cell
		vO *x.Opta // option cell - found in Cells, for DoCoverOthers & UnCoverOthers
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
	v, vO = i, &l.optaS[i]
	l.SetItem(i, v, vO) // Remember main
	if vO.Next == i {   // => X7 - deadend
		if on.Fail != nil {
			on.Fail.Do()
		} // we have a Dead end!
		goto X7 // >>>>>>>>>>
	}
	d.DoCover(i)                       // Inline ========================================
	v, vO = vO.Next, &l.optaS[vO.Next] // first option under i
X5: // [Try x[l].] ===========================================================
	if v == i { // => X7 - no more
		goto X7 // >>>>>>>>>>
	}
	d.DoCoverOthers(v)  // Inline ========================================
	l.SetItem(i, v, vO) // Remember cell
	l.Level++           // Incr Level
	goto X2             // >>>>>>>>>>
X6: // [Try again] ===========================================================
	v, vO = l.GetBoth()                            // Restore cell
	d.UnCoverOthers(v)                             // Inline ========================================
	i, v, vO = vO.Root, vO.Next, &l.optaS[vO.Next] // next option under i
	goto X5                                        // >>>>>>>>>>
X7: // [Backtrack] ===========================================================
	d.UnCover(i) // Inline ========================================
X8: // [Leave level] =========================================================
	if l.Level == 0 { // => X9
		goto X9 // >>>>>>>>>>
	}
	l.Level-- // Decr Level
	goto X6   // >>>>>>>>>>

X9: // we are done
	return
}
