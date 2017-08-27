package math

import (
	"math"
	"testing"
)

const epsilon float64 = 1e-10

func testGammaLn(x float64, t *testing.T) {
	expected := math.Log(math.Gamma(x))
	actual := GammaLn(x)
	if math.Abs(expected-actual) > epsilon {
		t.Errorf("GammaLn(%f) = %f, expected %f", x, actual, expected)
	}
}

func TestGammaLn(t *testing.T) {
	testGammaLn(0.0, t)
	testGammaLn(0.5, t)
	testGammaLn(1.0, t)
	testGammaLn(2.0, t)
	testGammaLn(5.0, t)
}
