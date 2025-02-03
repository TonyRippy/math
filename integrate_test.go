package math

import (
	"math"
	"testing"
)

const (
	EPS = 1e-10
)

func TestRombergWithMidPoint(t *testing.T) {
	// \int(1) = x
	f := func(x float64) float64 {
		return 1
	}
	actual, err := Romberg(MidPoint(f, 1, 5), EPS)
	if err != nil {
		t.Error(err)
		return
	}
	expected := 5.0 - 1.0
	if math.Abs(actual-expected) > EPS {
		t.Errorf("Romberg: Expected %f, got %f.", expected, actual)
	}

	// \int(x) = x^2 / 2
	f = func(x float64) float64 {
		return x
	}
	actual, err = Romberg(MidPoint(f, 2, 4), EPS)
	if err != nil {
		t.Error(err)
		return
	}
	expected = (4*4 - 2*2) / 2.0
	if math.Abs(actual-expected) > EPS {
		t.Errorf("Romberg: Expected %f, got %f.", expected, actual)
	}

	// \int(6*x^2) = 2x^3
	f = func(x float64) float64 {
		return x * x * 6.0
	}
	actual, err = Romberg(MidPoint(f, 0, 4), EPS)
	if err != nil {
		t.Error(err)
		return
	}
	expected = 128.0
	if math.Abs(actual-expected) > EPS {
		t.Errorf("Romberg: Expected %f, got %f.", expected, actual)
	}
}

func TestRombergWithSqrtLower(t *testing.T) {
	// \int(1/x) = ln(x)
	f := func(x float64) float64 {
		return 1 / x
	}
	actual, err := Romberg(MidPointSqrtLower(f, 0.01, 4), EPS)
	if err != nil {
		t.Error(err)
		return
	}
	expected := math.Log(4) - math.Log(0.01)
	if math.Abs(actual-expected) > EPS {
		t.Errorf("Romberg: Expected %f, got %f.", expected, actual)
	}
}

func TestRungeKutta(t *testing.T) {
	// y[0]: f(x)  = x
	// y[1]: f'(x) = 1
	//       f"(x) = 0
	ode := func(x float64, y []float64, dydx []float64) {
		dydx[0] = 1
		dydx[1] = 0
	}
	y := []float64{1, 1} // values at x=1
	if err := RungeKutta(ode, 1, 5, y, 1, EPS); err != nil {
		t.Error(err)
		return
	} else {
		// What are the values at x=5?
		expected := []float64{5, 1}
		for i, actual := range y {
			if math.Abs(actual-expected[i]) > EPS {
				t.Errorf("RungeKutta, y[%d]: Expected %f, got %f.", i, expected[i], actual)
			}
		}
	}

	// y[0]: f(x)  = x^2
	// y[1]: f'(x) = 2x
	//       f"(x) = 2
	ode = func(x float64, y []float64, dydx []float64) {
		dydx[0] = x * 2.0
		dydx[1] = 2
	}
	y = []float64{4, 4} // Values at x=2
	if err := RungeKutta(ode, 2, 4, y, 2, EPS); err != nil {
		t.Error(err)
	} else {
		// What is the value at x=4?
		expected := []float64{16, 8}
		for i, actual := range y {
			if math.Abs(actual-expected[i]) > EPS {
				t.Errorf("RungeKutta, y[%d]: Expected %f, got %f.", i, expected[i], actual)
			}
		}
	}

	// y[0]: f(x)  = 2x^3-4x+1
	// y[1]: f'(x) = 6x^2-4
	// y[2]: f"(x) = 12x
	//             = 12
	ode = func(x float64, y []float64, dydx []float64) {
		dydx[0] = 6.0*x*x - 4.0
		dydx[1] = 12.0 * x
		dydx[2] = 12.0
	}
	y = []float64{-1, 2, 12} // Values at x=1
	if err := RungeKutta(ode, 1, 3, y, 0.5, EPS); err != nil {
		t.Error(err)
	} else {
		// What is the value at x=3?
		expected := []float64{2*3*3*3 - 4*3 + 1, 6*3*3 - 4, 12 * 3}
		for i, actual := range y {
			if math.Abs(actual-expected[i]) > EPS {
				t.Errorf("RungeKutta, y[%d]: Expected %f, got %f.", i, expected[i], actual)
			}
		}
	}
}
