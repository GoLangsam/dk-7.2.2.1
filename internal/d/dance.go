// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d

// ===========================================================================

func (a *D) Dance(c Index) {

	var v Index // some option under c - think: vertical

	{
		{
			a.DoCover(c) // Inline ========================================
		}
	}

	for v = a.OptaS[c].Next; v != c; v = a.OptaS[v].Next {

		a.Stack = append(a.Stack, v) // a.Stack.Push(v)

		a.DoCoverOthers(v) // Inline ========================================

		a.On.Next()

		a.UnCoverOthers(v) // Inline ========================================

		a.Stack = a.Stack[:len(a.Stack)-1] // a.Stack.Drop()

	}

	{
		{
			a.UnCover(c) // Inline ========================================
		}
	}

}

// ===========================================================================
