package math

/* Given N, calculate N! */
func Factorial(n uint) uint {
	var f uint = 1
	for ; n > 1; n-- {
		f *= n
	}
	return f
}
