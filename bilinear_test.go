package math

import (
	"math"
	"testing"
)

func mostlyEqual(expected, actual float64) bool {
	const epsilon float64 = 1e-10
	return math.Abs(expected-actual) < epsilon
}

func testF(f func(float64, float64) float64, x float64, y float64, expected float64, t *testing.T) {
	actual := f(x, y)
	if !mostlyEqual(expected, actual) {
		t.Errorf("f(%f, %f) = %f, expected %f", x, y, actual, expected)
	}
}

func TestBilinear8(t *testing.T) {
	f := bilinear8(0, 1, 0, 1, 0, 1, 1, 1)
	testF(f, 0.0, 0.0, 0.0, t)
	testF(f, 0.0, 0.5, 0.5, t)
	testF(f, 0.5, 0.0, 0.5, t)
	testF(f, 0.0, 1.0, 1.0, t)
	testF(f, 1.0, 0.0, 1.0, t)
	testF(f, 1.0, 1.0, 1.0, t)
	testF(f, 0.5, 0.5, 0.75, t)
}

func TestBilinear4(t *testing.T) {
	f := bilinear4(0, 1, 0, 1, 0, 1, 1, 1)
	testF(f, 0.0, 0.0, 0.0, t)
	testF(f, 0.0, 0.5, 0.5, t)
	testF(f, 0.5, 0.0, 0.5, t)
	testF(f, 0.0, 1.0, 1.0, t)
	testF(f, 1.0, 0.0, 1.0, t)
	testF(f, 1.0, 1.0, 1.0, t)
	testF(f, 0.5, 0.5, 0.75, t)
}
