// Methods for performing linear interpolation.

package math

/*
SolveLinear calculates the coefficients of a line given two points.

A line is often expressed as the following formula: y = Mx + B.
This function, given two (x,y) points, will calcuate and return the values for M, B.
*/
func SolveLinear(x1, y1, x2, y2 float64) (M, B float64) {
	M = (y2 - y1) / (x2 - x1)
	B = y1 - M*x1
	return
}

/*
LinearFunction returns a function that performs linear interpolation.

Given the values of two (x,y) points, this will return a function f(x)
that allows one to interpolate values between the two given points.
*/
func LinearFunction(x1, y1, x2, y2 float64) func(float64) float64 {
	M, B := SolveLinear(x1, y1, x2, y2)
	return func(x float64) float64 {
		return M*x + B
	}
}
