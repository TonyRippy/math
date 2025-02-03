/*
Methods for performing numerical integration, based on the methods described in:
Numerical Recipes in C, Second Edition, by Press, Teukolsky, Vetterling, Flannery.

Some modifications to the original:
  - Switched from 1-based arrays to 0-based arrays.
  - Removed use of global variables.
  - Removed the storage of intermediate state.
*/

package math

import (
	"errors"
	"math"
)

// Type signature of a stepper routine that can be used with Romberg integration, below.
type RombergStepper func(int, float64) float64

/*
This routine creates a stepper function that computes the nth stage of
refinement of an extended midpoint rule. f is the function to be
integrated between limits a and b, also provided.  When called with
n=1, the routine returns the crudest estimate of \int_a^b
f(x)dx. Subsequent calls with n=2,3,... (in that sequential order)
will improve the accuracy of s by adding (2/3)*3^{n-1} additional
interior points.
*/
func MidPoint(f func(float64) float64, a float64, b float64) RombergStepper {
	return func(n int, s float64) float64 {
		if n == 1 {
			return (b - a) * f(0.5*(a+b))
		}
		it := 1
		for j := 1; j < n-1; j++ {
			it *= 3
		}
		tnm := float64(it)
		del := (b - a) / (3.0 * tnm)
		ddel := del + del
		// The added points alternate in spacing between del and ddel.
		x := a + 0.5*del
		sum := 0.0
		for j := 1; j <= it; j++ {
			sum += f(x)
			x += ddel
			sum += f(x)
			x += del
		}
		// The new sum is combined with the old integral to give a refined integral.
		return (s + (b-a)*sum/tnm) / 3.0
	}
}

/*
This routine is an exact replacement for MidPoint, i.e., returns the
nth stage of refinement of the integral of f from a to b, except that
the function is evaluated at evenly spaced points in 1/x rather than
in x.  This allows the upper limit b to be as large and positive as
the computer allows, or the lower limit a to be as large and
negative, but not both. a and b must have the same sign.
*/
func MidPointInf(f func(float64) float64, a float64, b float64) RombergStepper {
	// Change the limits of integration.
	ff := func(x float64) float64 {
		return f(1.0/x) / (x * x)
	}
	aa := 1.0 / b
	bb := 1.0 / a
	return MidPoint(ff, aa, bb)
}

/*
Romberg integration on an open interval, with desired accuracy eps.
Normally stepper will be an open formula, not evaluating its function at the
endpoints. It is assumed that stepper triples the number of steps on each call,
and that its error series contains only even powers of the number of steps.  The
routines MidPoint and MidPointInf are possible choices.
(Also midsql, midsqu and midexp, which are not implemented here.)
*/
func Romberg(stepper RombergStepper, eps float64) (ss float64, err error) {
	const (
		JMAX  = 14
		JMAXP = JMAX + 1
		K     = 5
	)
	ss = 0.0
	h := make([]float64, JMAXP+1)
	h[1] = 1.0
	s := make([]float64, JMAXP)
	for j := 1; j <= JMAX; j++ {
		s[j] = stepper(j, s[j-1])
		if j >= K {
			i := j - K + 1
			var dss float64
			ss, dss, err = polint(h[i:], s[i:], K, 0.0)
			if err != nil {
				return ss, err
			}
			if math.Abs(dss) <= eps*math.Abs(ss) {
				return ss, nil
			}
		}
		// This is where the assumption of step tripling and even error series is used.
		h[j+1] = h[j] / 9.0
	}
	return ss, errors.New("Too many steps in Romberg")
}

/*
A user-supplied function that evaluates dydx at x, given the dependent variables y[].
*/
type OdeFunc func(x float64, y []float64, dydx []float64)

/*
Given values for n variables y[] and their derivatives dydx[] known at x, use
the fifth-order Cash-Karp Runge-Kutta method to advance the solution over an
interval h and return the incremented variables as yout[].  Also return an
estimate of the local truncation error in yerr[] using the embedded fourth-order
method. The user supplies the routine derivs(x, y, dxdy), which returns
derivatives dydx at x.
*/
func rkck(y []float64, dydx []float64, n int, x float64, h float64, yout []float64, yerr []float64, derivs OdeFunc) {
	const (
		a2  = 0.2
		a3  = 0.3
		a4  = 0.6
		a5  = 1.0
		a6  = 0.875
		b21 = 0.2
		b31 = 3.0 / 40.0
		b32 = 9.0 / 40.0
		b41 = 0.3
		b42 = -0.9
		b43 = 1.2
		b51 = -11.0 / 54.0
		b52 = 2.5
		b53 = -70.0 / 27.0
		b54 = 35.0 / 27.0
		b61 = 1631.0 / 55296.0
		b62 = 175.0 / 512.0
		b63 = 575.0 / 13824.0
		b64 = 44275.0 / 110592.0
		b65 = 253.0 / 4096.0
		c1  = 37.0 / 378.0
		c3  = 250.0 / 621.0
		c4  = 125.0 / 594.0
		c6  = 512.0 / 1771.0
		dc1 = c1 - 2825.0/27648.0
		dc3 = c3 - 18575.0/48384.0
		dc4 = c4 - 13525.0/55296.0
		dc5 = -277.0 / 14336.0
		dc6 = c6 - 0.25
	)
	ak2 := make([]float64, n)
	ak3 := make([]float64, n)
	ak4 := make([]float64, n)
	ak5 := make([]float64, n)
	ak6 := make([]float64, n)
	ytemp := make([]float64, n)
	// First step.
	for i := 0; i < n; i++ {
		ytemp[i] = y[i] + h*(b21*dydx[i])
	}
	// Second step.
	derivs(x+a2*h, ytemp, ak2)
	for i := 0; i < n; i++ {
		ytemp[i] = y[i] + h*(b31*dydx[i]+b32*ak2[i])
	}
	// Third step.
	derivs(x+a3*h, ytemp, ak3)
	for i := 0; i < n; i++ {
		ytemp[i] = y[i] + h*(b41*dydx[i]+b42*ak2[i]+b43*ak3[i])
	}
	// Fourth step.
	derivs(x+a4*h, ytemp, ak4)
	for i := 0; i < n; i++ {
		ytemp[i] = y[i] + h*(b51*dydx[i]+b52*ak2[i]+b53*ak3[i]+b54*ak4[i])
	}
	// Fifth step.
	derivs(x+a5*h, ytemp, ak5)
	for i := 0; i < n; i++ {
		ytemp[i] = y[i] + h*(b61*dydx[i]+b62*ak2[i]+b63*ak3[i]+b64*ak4[i]+b65*ak5[i])
	}
	// Sixth step.
	derivs(x+a6*h, ytemp, ak6)
	for i := 0; i < n; i++ {
		// Accumulate increments with proper weights.
		yout[i] = y[i] + h*(c1*dydx[i]+c3*ak3[i]+c4*ak4[i]+c6*ak6[i])
		// Estimate error as difference between fourth and fifth order methods.
		yerr[i] = h * (dc1*dydx[i] + dc3*ak3[i] + dc4*ak4[i] + dc5*ak5[i] + dc6*ak6[i])
	}
}

// Type signature of a stepper routine that can be used with the Runge-Kutta driver, below.
type rungeKuttaStepper func([]float64, []float64, int, float64, float64, float64, []float64, OdeFunc) (float64, float64, error)

/*
Fifth-order Runge-Kutta step with monitoring of local truncation error to ensure
accuracy and adjust step size.  Inputs are the dependent variable vector y[] and
its derivative dydx[] at the starting value of the independent variable x.  Also
input are the step size to be attempted h, the required accuracy eps, and the
vector yscal[] against which the error is scaled.  On output, the y vector is
replaced by their new values, and the function returns new values for x and h
respectively.
*/
func rkqs(y []float64, dydx []float64, n int, x float64, h float64, eps float64, yscal []float64, derivs OdeFunc) (float64, float64, error) {
	const (
		SAFETY  = 0.9
		PGROW   = -0.2
		PSHRINK = -0.25
		ERRCON  = 1.89e-4 // ~= math.Pow(5/SAFETY, 1/PGROW)
	)
	var errmax float64
	ytemp := make([]float64, n)
	yerr := make([]float64, n)
	for {
		// Take a step.
		rkck(y, dydx, n, x, h, ytemp, yerr, derivs)
		// Evaluate accuracy.
		errmax = 0.0
		for i := 0; i < n; i++ {
			errmax = math.Max(errmax, math.Abs(yerr[i]/yscal[i]))
		}
		errmax /= eps // Scale relative to required tolerance.
		if errmax <= 1.0 {
			break // Step succeeded.
		}
		// Truncation error too large, reduce step size.
		htemp := SAFETY * h * math.Pow(errmax, PSHRINK)
		if h >= 0.0 {
			h = math.Max(htemp, 0.1*h)
		} else {
			h = math.Min(htemp, 0.1*h)
		}
		// No more than a factor of 10.
		xnew := x + h
		if xnew == x {
			return xnew, h, errors.New("step size underflow in rkqs")
		}
	}

	// Compute size of next step.
	var hnext float64
	if errmax > ERRCON {
		hnext = SAFETY * h * math.Pow(errmax, PGROW)
	} else {
		// No more than a factor of 5 increase.
		hnext = 5.0 * h
	}
	xnext := x + h
	for i := 0; i < n; i++ {
		y[i] = ytemp[i]
	}
	return xnext, hnext, nil
}

/*
A Runge-Kutta driver with adaptive step size control. Integrate starting values
ystart[] from x1 to x2 with accuracy eps, storing intermediate results in global
variables. (???) h1 should be set as a guessed first step size, hmin as the
minimum allowed step size (can be zero).  On output ystart[] is replaced by
values at the end of the integration interval.  derivs is the user-supplied
routine for calculating the right-hand side derivative, while step is the
stepper routine to be used.
*/
func odeint(ystart []float64, nvar int, x1 float64, x2 float64, eps float64, h1 float64, hmin float64, derivs OdeFunc, stepper rungeKuttaStepper) error {
	const (
		MAXSTP int     = 10000 // TODO: Make this a flag?
		TINY   float64 = 1.0e-30
	)
	yscal := make([]float64, nvar)
	y := make([]float64, nvar)
	dydx := make([]float64, nvar)
	x := x1
	h := math.Copysign(h1, x2-x1)
	for i := 0; i < nvar; i++ {
		y[i] = ystart[i]
	}
	for nstp := 0; nstp < MAXSTP; nstp++ { // Take at most MAXSTP steps.
		derivs(x, y, dydx)
		for i := 0; i < nvar; i++ {
			// Scaling used to monitor accuracy. This general-purpose choice can be
			// modified if need be.
			yscal[i] = math.Abs(y[i]) + math.Abs(dydx[i]*h) + TINY
		}
		if (x+h-x2)*(x+h-x1) > 0.0 {
			// If step size can overshoot, decrease.
			h = x2 - x
		}
		var err error
		x, h, err = stepper(y, dydx, nvar, x, h, eps, yscal, derivs)
		if err != nil {
			return err
		}
		if (x-x2)*(x2-x1) >= 0.0 { // Are we done?
			for i := 0; i < nvar; i++ {
				ystart[i] = y[i]
			}
			return nil // Normal exit.
		}
		if math.Abs(h) <= hmin {
			return errors.New("step size too small in odeint")
		}
	}
	return errors.New("too many steps in odeint")
}

/*
A Runge-Kutta driver with adaptive step size control.

Integrates the set of ordinary differential equations specified by ode from x1
to x2, advancing the solution in steps, using starting values y[] known at x1.
The initial step size is h, but the driver will adjust the step size trying to
achieve the targeted accuracy eps. The final values at x2 are returned in y[].
*/
func RungeKutta(ode OdeFunc, x1 float64, x2 float64, y []float64, h float64, eps float64) error {
	return odeint(y, len(y), x1, x2, eps, h, 0.0, ode, rkqs)
}
