package math

import (
	"testing"
)

// TestLinear checks the linear interpolation code against known good values.
func TestLinear(t *testing.T) {
	type point struct {
		X, Y float64
	}
	cases := []struct {
		Name           string
		X1, Y1, X2, Y2 float64
		Want           []point
	}{
		{"StartAtZero", 0, 0, 1, 1, []point{{0.5, 0.5}}},
		{"EndAtZero", -1, 1, 0, 0, []point{{-0.5, 0.5}}},
		{"SteepSlope", 1, 3, 2, 5001, []point{{1.1, 502.8}, {1.5, 2502}}},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			f := LinearFunction(tc.X1, tc.Y1, tc.X2, tc.Y2)
			if !mostlyEqual(f(tc.X1), tc.Y1) {
				t.Errorf("f(%f): got %f, want %f", tc.X1, f(tc.X1), tc.Y1)
			}
			if !mostlyEqual(f(tc.X2), tc.Y2) {
				t.Errorf("f(%f): got %f, want %f", tc.X2, f(tc.X2), tc.Y2)
			}
			for _, want := range tc.Want {
				if !mostlyEqual(f(want.X), want.Y) {
					t.Errorf("f(%f): got %f, want %f", want.X, f(want.X), want.Y)
				}
			}
		})
	}
}
