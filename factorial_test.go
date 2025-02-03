package math

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	cases := [][]uint{
		[]uint{0, 1},
		[]uint{1, 1},
		[]uint{2, 2},
		[]uint{5, 5 * 4 * 3 * 2 * 1},
	}
	for _, c := range cases {
		actual := Factorial(c[0])
		expected := c[1]
		if actual != expected {
			t.Errorf("Factorial(%d): Expected %d, got %d.", c[0], expected, actual)
		}
	}
}
