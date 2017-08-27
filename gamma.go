package math

import (
	"math"
)

// Gamma functions, ripped from the pages of:
// Numerical Recipies in C, Second Edition, by Press, Teukolsky, Vetterling, Flannery.

func GammaLn(xx float64) float64 {
	var cof = [6]float64{76.18009172947146, -86.50532032941677, 24.01409824083091, -1.231739572450155, 0.1208650973866179e-2, -0.5395239384953e-5}
	y := xx
	x := xx
	tmp := x + 5.5
	tmp -= (x + 0.5) * math.Log(tmp)
	ser := 1.000000000190015
	for j := 0; j <= 5; j += 1 {
		y += 1
		ser += cof[j] / y
	}
	return -tmp + math.Log(2.5066282746310005*ser/x)
}
