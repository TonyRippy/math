// Methods for performing cubic interpolation. 

package math

/*
  Returns a closure that performs cubic interpolation.
  This returns a somewhat inefficient variant that peforms
  all computation at runtime and captures all 6 input values.
  Useful for testing purposes.
*/
func cubic6(x1, y1, dy1, x2, y2, dy2 float64) func (float64) float64 {
	// For implementation details, see:
  // https://en.wikipedia.org/wiki/Cubic_Hermite_spline
	return func(x float64) float64 {
		dx := x2 - x1
		t := (x - x1) / dx
		tt := t * t
		ttt := tt * t
		return (
			(2*ttt - 3*tt + 1) * y1 + 
			(ttt - 2*tt + t) * dy1 * dx +
			(-2*ttt + 3*tt) * y2 +
			(ttt - tt) * dy2 * dx)
	}
}

/*
  A cubic spline can be expressed as a third-degree polynomial function:
    f(x) = Ax^3 + Bx^2 + Cx + D
  
  This function, given the values of x, f(x) and f'(x) at two end points,
  will calcuate and return the values for A, B, C, D.     
*/
func SolveCubic(x1, y1, dy1, x2, y2, dy2 float64) (A, B, C, D float64) {
	// Yeah, sorry about this mess... I will try and post a derivation for this.
	dx := x2 - x1
	xx1 := x1 * x1
	xx2 := x2 * x2
	
	a1 := (xx2 * x2) - (xx1 * x1)
	a2 := a1
	a1 -= 3 * xx1 * dx
	a2 -= 3 * xx2 * dx

	b1 := xx2 - xx1
	b2 := b1
	b1 -= (2 * x1 * dx)
	b2 -= (2 * x2 * dx)

	e1 := y2 - y1
	e2 := e1
	e1 -= dy1 * dx
	e2 -= dy2 * dx

	f := b1 / b2

	A = (e1 - f * e2) / (a1 - f * a2)
	B = (e1 - A * a1) / b1
	C = dy1 - (3 * A * xx1) - (2 * B * x1)
	D = y1 - (A * xx1 * x1) - (B * xx1) - (C * x1)
	return
}

/*
  Returns a closure that performs cubic interpolation.
  This returns a more efficient closure that is expensive to create,
  but minimizes runtime computation. 
  Only requires the capture of 4 float64 values.
*/
func cubic4(x1, y1, dy1, x2, y2, dy2 float64) func (float64) float64 { 
	A, B, C, D := SolveCubic(x1, y1, dy1, x2, y2, dy2)
	return func(x float64) float64 {
		xx := x * x
		return A * xx * x + B * xx + C * x + D
  }
}

/*
  CubicFunction returns a function that performs cubic interpolation.

  Given the values of x, f(x) and f'(x) at two end points, this will return
  a function f(x) that allows one to interpolate values between the two
  given points.
*/
func CubicFunction(x1, y1, dy1, x2, y2, dy2 float64) func (float64) float64 { 
	return cubic4(x1, y1, dy1, x2, y2, dy2)
}
