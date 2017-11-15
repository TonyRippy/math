package math

import (
	"fmt"
)

/*
Calculates the number of ways that k items can be chosen from a list of n items.
A.k.a. "N choose K".
Given by the formula: n! / (k! * (n-k)!)
Attempts to be clever about the calculation to minimize the risk of overflow.
*/
func Combinations(n uint, k uint) uint {
	if k > n {
		return 0
	}
	j := n - k
	if j > k {
		k, j = j, k
	}
	if k == n {
		return 1
	}
	result := n
	var hi uint = n - 1
	if hi > k {
		result *= hi
		hi -= 1
		if hi > k {
			result *= hi
			hi -= 1
		}
	}
	var lo uint = 2
	for ; (lo <= j) && (result % lo == 0); lo++ {
			result /= lo
	}
	for ; hi > k; hi-- {
		result *= hi
	}
	for ; lo <= j; lo++ {
		result /= lo
	}
	return result
}

/*
Generates k-combinations of a collection of n items.

Enumerates all the ways that k items can be chosen from a list of n items.
Calls a provided function f multiple times, each call with a different
permutation. The argument to f is a slice of boolean values that indicate which
of the n elements are part of that permutation (true) and which aren't. (false)

The callback f should not modify the slice. Modifying the slice will result
in undefined behavior.

This method does not depend on recursion and uses O(n) memory, so is suitable
for enumerating potentially large numbers of combinations.
*/
func WalkCombinations(n int, k int, f func([]bool)) error {
	if n < 0 {
		return fmt.Errorf("Cannot choose from a negative number of items. (%d)", n)
	}	
	if k > n {
		return fmt.Errorf("Cannot choose %d out of %d", k, n)
	}
	// Create an array to track which items are part of the current
	// permutation. For a given item, we always try 'false' first,
	// then 'true'. That way we can avoid recursion which could
	// overflow the stack when n is large.
	perm := make([]bool, n)
	if n == 0 {
		f(perm)
		return nil
	}
	// Set up the inital permutation, where the last elements are chosen.
	// [0 1 2 3 4 5], k=2, n=6 --> [f, f, f, f, t, t]
	// [0 1 2 3 4], k=3, n=5 : --> [f, f, t, t, t]
	x := n - k
	for i, _ := range(perm) {
		perm[i] = (i >= x)
	}
	pop := false
	remain := 0
	i := n
	for ; i >= 0; {
		if i == n {
			// Permutation fully specified. Call the function:
			f(perm)
			// ... then pop the stack.
			pop = true
			i -= 1
			continue
		}
		if perm[i] {
			// The element we're considering has already been true,
			// so we need to pop the stack.
			pop = true
			perm[i] = false  // unpick the value
			remain += 1  // one more to pick
			i -= 1  // move back to previous item
			continue
		}
		// If you get here, perm[i] == false.
		if !pop {
			// We haven't tried "false" yet. Can we?
			if n - i > remain {
				// We can still say "false", because there are enough elements
				// left to pick the right number of "true"s.
				// Leave this element false and move forward.
				i += 1
				continue
			}
		}
		// Can we say "true"?
		if remain > 0 {
			// Try it.
			perm[i] = true
			remain -= 1
			pop = false
			i += 1
			continue
		}
		// Otherwise pop the stack.
		pop = true
		i -= 1
	}
	// Done. We've generated all the permutations.
	return nil
}
