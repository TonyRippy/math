package math

import (
	"testing"
)

// TestCubic4vs6 tests the cubic code by comparing the output of two different
// interpolation methods that should yield the same numerical result.
func TestCubic4vs6(t *testing.T) {
	cases := []struct{
		Name string
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
