// Methods for performing bilinear interpolation. 
// https://en.wikipedia.org/wiki/Bilinear_interpolation

package math

/*
  Returns a closure that performs the interpolation.
  This returns a somewhat inefficient variant that peforms
  all computation at runtime and captures all 8 input values.
  Useful for testing purposes.
*/
func bilinear8(x1, x2, y1, y2, q11, q12, q21, q22 float64) func (float64, float64) float64 { 
	return func(x float64, y float64) float64 {
		return (
			((q11 * (x2 - x) * (y2 - y)) +
				(q21 * (x - x1) * (y2 - y)) +
				(q12 * (x2 - x) * (y - y1)) +
				(q22 * (x - x1) * (y - y1))) /
				((x2 - x1) * (y2 - y1)))
	}
}

/*
  Bilinear interpolation can be expressed as a linear transorm:
    f(x,y) = A * x * y + B * x + C * y + D 
  
  This function, given the corners of a rectangular region:
    q11 = f(x1, y1)
    q12 = f(x1, y2)
    q21 = f(x2, y1)
    q22 = f(x2, y2)
  will calcuate and return the values for A, B, C, D.     
*/
func SolveBilinear(x1, x2, y1, y2, q11, q12, q21, q22 float64) (A, B, C, D float64) {
	A = (q11 - q12 - q21 + q22) / (x1*y1 - x1*y2 - x2*y1 + x2*y2)
	B = (q21 - q11 + A*(x1*y1 - x2*y1)) / (x2 - x1)
	C = (q12 - q11 + A*(x1*y1 - x1*y2)) / (y2 - y1)
	D = (q11 - A*x1*y1 - B*x1 - C*y1)
	return
}

/*
  Returns a closure that performs bilinear interpolation.
  This returns a more efficient closure that is more expensive
  to create, but minimizes runtime computation.
  Only requires the capture of 4 float64 values.
*/
func bilinear4(x1, x2, y1, y2, q11, q12, q21, q22 float64) func (float64, float64) float64 { 
	A, B, C, D := SolveBilinear(x1, x2, y1, y2, q11, q12, q21, q22)
	return func(x float64, y float64) float64 {
		return A * x * y + B * x + C * y + D
  }
}

/*
  Returns a function that can perform bilinear interpolation. 

  Given a bounding rectangle (x1,y1)-(x2,y2) and values at the corners:
    q11 = f(x1, y1)
    q12 = f(x1, y2)
    q21 = f(x2, y1)
    q22 = f(x2, y2)
  This will return the function f(x,y).
*/
func BilinearFunction(x1, x2, y1, y2, q11, q12, q21, q22 float64) func (float64, float64) float64 { 
	return bilinear4(x1, x2, y1, y2, q11, q12, q21, q22)
}
