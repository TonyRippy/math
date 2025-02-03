// Methods for performing cubic interpolation.

package math

import (
	"math"
)

/*
Returns a closure that performs cubic interpolation.
This returns a somewhat inefficient variant that peforms
all computation at runtime and captures all 6 input values.
Useful for testing purposes.
*/
func cubic6(x1, y1, dy1, x2, y2, dy2 float64) func(float64) float64 {
	// For implementation details, see:
	// https://en.wikipedia.org/wiki/Cubic_Hermite_spline
	return func(x float64) float64 {
		dx := x2 - x1
		t := (x - x1) / dx
		tt := t * t
		ttt := tt * t
		return ((2*ttt-3*tt+1)*y1 +
			(ttt-2*tt+t)*dy1*dx +
			(-2*ttt+3*tt)*y2 +
			(ttt-tt)*dy2*dx)
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

	A = (e1 - f*e2) / (a1 - f*a2)
	B = (e1 - A*a1) / b1
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
func cubic4(x1, y1, dy1, x2, y2, dy2 float64) func(float64) float64 {
	A, B, C, D := SolveCubic(x1, y1, dy1, x2, y2, dy2)
	return func(x float64) float64 {
		return ((A*x+B)*x+C)*x + D
	}
}

/*
CubicFunction returns a function that performs cubic interpolation.

Given the values of x, f(x) and f'(x) at two end points, this will return
a function f(x) that allows one to interpolate values between the two
given points.
*/
func CubicFunction(x1, y1, dy1, x2, y2, dy2 float64) func(float64) float64 {
	return cubic4(x1, y1, dy1, x2, y2, dy2)
}

/*
FritschCarlsonTangents calculates tangents for a set of points
that ensure monotonicity for a resulting Hermite spline.

This function makes some important assumptions:
 1. That the input arrays have the same length.
 2. That the data points are monotonic.
 3. That the input points are sorted on the x axis, ascending.

These assumptions are not verified by the method.
*/
func FritschCarlsonTangents(xs, ys []float64) []float64 {
	// For implementation details, see:
	// https://en.wikipedia.org/wiki/Monotone_cubic_interpolation
	n := len(xs)
	if n == 0 {
		return []float64{}
	}
	if n == 1 {
		return []float64{0}
	}
	// Compute the slopes of the secant lines between successive points
	d := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		d[i] = (ys[i+1] - ys[i]) / (xs[i+1] - xs[i])
	}
	// Compute provisional tangents
	m := make([]float64, n)
	m[0] = d[0]
	m[n-1] = d[n-2]
	for i := 1; i < n-1; i++ {
		if d[i] == 0.0 {
			m[i] = 0.0
			i += 1
			m[i] = 0.0
			continue
		}
		if math.Signbit(d[i-1]) != math.Signbit(d[i]) {
			m[i] = 0.0
		} else {
			m[i] = (d[i-1] + d[i]) / 2
		}
	}
	// Adjust tangents to keep monoticity.
	for i := 0; i < n-1; i++ {
		dk := d[i]
		ak := m[i] / dk
		bk := m[i+1] / dk
		sqsum := ak*ak + bk*bk
		if sqsum > 9.0 {
			tk := 3.0 / math.Sqrt(sqsum)
			m[i] = tk * ak * dk
			m[i+1] = tk * bk * dk
		}
	}
	return m
}
