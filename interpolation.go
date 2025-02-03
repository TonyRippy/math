/*
Methods for performing interpolation, ported from the original implementation from:
Numerical Recipies in C, Second Edition, by Press, Teukolsky, Vetterling, Flannery.

Some modifications to the original:
  - Switched from 1-based arrays to 0-based arrays.
*/

package math

import (
	"errors"
	"math"
)

/*
Given arrays xa[] and ya[], and given a value x, this routine returns
a value y, and an error estimate dy. If P(x) is the polynomial of
degree n-1 such that P(xa_i) = ya_i, then the returned value is y = P(x).
*/
func polint(xa []float64, ya []float64, n int, x float64) (y float64, dy float64, err error) {
	err = nil
	c := make([]float64, n)
	d := make([]float64, n)
	// Here we find the index ns of the closest table entry,
	var ns int
	dif := math.MaxFloat64
	for i := 0; i < n; i++ {
		dift := math.Abs(x - xa[i])
		if dift < dif {
			ns = i
			dif = dift
		}
		// and initialize the tableau of c's and d's.
		c[i] = ya[i]
		d[i] = ya[i]
	}
	y = ya[ns] // This is the initial approximation to y.
	ns -= 1
	for m := 1; m < n; m++ { // For each column of the tableau,
		for i := 0; i < n-m; i++ { // we loop over the current c's and d's and update them.
			ho := xa[i] - x
			hp := xa[i+m] - x
			w := c[i+1] - d[i]
			den := ho - hp
			if den == 0.0 {
				err = errors.New("Error in routine polint")
				return
			}
			den = w / den
			d[i] = hp * den // Here the c's and d's are updated.
			c[i] = ho * den
		}
		if 2*ns < n-m {
			dy = c[ns+1]
		} else {
			dy = d[ns]
			ns -= 1
		}
		y += dy
		// After each column of the tableau is completed, we decide which
		// correction, c or d, we want to add to our accumulating value of
		// y, i.e., which path to take through the tableau -- forking up
		// or down.  We do this in such a way as to take the most
		// "straight line" route through the tableau to its apex, updating
		// ns accordingly to keep track of where we are. This route keeps
		// the partial approximations centered (insofar as possible) on
		// the target x. The last dy added is thus the error indication.
	}
	return
}
