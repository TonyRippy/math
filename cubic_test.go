package math

import (
	"testing"
)

// TestCubic4vs6 tests the cubic code by comparing the output of two different
// interpolation methods that should yield the same numerical result.
func TestCubic4vs6(t *testing.T) {
	cases := []struct {
		Name                     string
		X1, Y1, DY1, X2, Y2, DY2 float64
	}{
		{"Line", 0, 1, 1, 1, 2, 1},
		{"CurveToFlat", 1, 0, 1, 10, 5, 0},
		{"CurveToSameSlope", 2, .333, .333, 5, .667, 0.333},
		{"Parabola", -1, 0, 1, 1, 0, -1},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			f4 := cubic4(tc.X1, tc.Y1, tc.DY1, tc.X2, tc.Y2, tc.DY2)
			if !mostlyEqual(f4(tc.X1), tc.Y1) {
				t.Errorf("f(%f) = %f, expected %f", tc.X1, f4(tc.X1), tc.Y1)
			}
			if !mostlyEqual(f4(tc.X2), tc.Y2) {
				t.Errorf("f(%f) = %f, expected %f", tc.X2, f4(tc.X2), tc.Y2)
			}
			f6 := cubic6(tc.X1, tc.Y1, tc.DY1, tc.X2, tc.Y2, tc.DY2)
			dx := (tc.X2 - tc.X1) / 100
			for x := tc.X1; x <= tc.X2; x += dx {
				expected := f6(x)
				actual := f4(x)
				if !mostlyEqual(actual, expected) {
					t.Errorf("f4(%f) = %f does not match f6(x) = %f", x, actual, expected)
				}
			}
		})
	}
}

func TestMonoticity(t *testing.T) {
	cases := []struct {
		Name   string
		Xs, Ys []float64
	}{
		{
			"Flat",
			[]float64{0, 10, 20, 30, 60},
			[]float64{3, 3, 3, 3, 3},
		},
		{
			"SmoothIncrease",
			[]float64{0, 1, 2, 3, 4},
			[]float64{2, 4, 6, 8, 10},
		},
		{
			"Curve",
			[]float64{0, 1, 2, 5, 60},
			[]float64{0, .1, .333, .667, .8},
		},
		{
			"SharpCurve",
			[]float64{0, 1, 2, 3, 4},
			[]float64{0, .1, .9, .95, .99},
		},
	}
	const pointsToCheck = 32
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			ms := FritschCarlsonTangents(tc.Xs, tc.Ys)
			n := len(ms)
			for i := 0; i < n-1; i++ {
				minx := tc.Xs[i]
				maxx := tc.Xs[i+1]
				f := cubic4(minx, tc.Ys[i], ms[i], maxx, tc.Ys[i+1], ms[i+1])
				dx := (maxx - minx) / pointsToCheck
				oy := tc.Ys[i]
				for x := minx + dx; x < maxx; x += dx {
					y := f(x)
					if y < oy {
						t.Errorf("Not monotonic: f(%f) = %f > f(%f) = %f", x-dx, oy, x, y)
						break
					}
					oy = y
				}
			}
		})
	}
}
