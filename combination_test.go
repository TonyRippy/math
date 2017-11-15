package math

import (
	"sort"
	"testing"
)

func TestCombinations(t *testing.T) {
	cases := [][]uint{
		[]uint{10, 0, 1},
		[]uint{10, 10, 1},
		[]uint{10, 1, 10},
		[]uint{10, 9, 10},
		[]uint{10, 3, 120},
		[]uint{10, 5, 252},
		[]uint{40, 10, 847660528},
		[]uint{90, 10, 5720645481903},
	}
	for _, c := range(cases) {
		expected := c[2]
		actual := Combinations(c[0], c[1])
		if actual != expected {
			t.Errorf("Combinations(%d, %d): Expected %d, got %d", c[0], c[1], expected, actual)
		}
	}
}

/*
A function that treats a boolean vector as a binary number,
converting it to an integer value.
*/
func permToInt(perm []bool) int {
	val := 0
	for i, v := range(perm) {
		if v {
			bit := 1 << uint(i)
			val = val | bit
		}
	}
	return val
}

func TestPermToInt(t *testing.T) {
	cases := [][]bool{
		[]bool{},
		[]bool{true},
		[]bool{false, true},
		[]bool{true, true},
		[]bool{false, false, true},
		[]bool{true, false, true},
		[]bool{false, true, true},
		[]bool{true, true, true},
	}
	for expected, perm := range(cases) {
		actual := permToInt(perm)
		if actual != expected {
			t.Errorf("PermToInt incorrect. Expected %d, got %d", expected, actual)
		}
	}
}

func TestChooseZero(t *testing.T) {
	// What happens if you choose zero elements?
	calls := 0
	WalkCombinations(0, 0, func(perm []bool) {
		calls += 1
		if len(perm) != 0 {
			t.Errorf("Expected empty permutation, got %v", perm)
		}
	})
	if calls != 1 {
		t.Errorf("WalkCombinations(0, 0): expected 1 callback, got %d", calls)
	}
	calls = 0
	WalkCombinations(100, 0, func(perm []bool) {
		calls += 1
		if len(perm) != 100 {
			t.Errorf("Expected len(perm) == 100, got %d", len(perm))
		}
		for _, v := range(perm) {
			if v {
				t.Errorf("Expected all items in permutation to be false! %v", perm)
			}
		}
	})
	if calls != 1 {
		t.Errorf("WalkCombinations(100, 0): expected 1 callback, got %d", calls)
	}	
}

func TestChooseOne(t *testing.T) {
	calls := make([]int, 0, 10)
	WalkCombinations(10, 1, func(perm []bool) {
		calls = append(calls, permToInt(perm))
	})
	expected := []int{
		1<<0,  // one bit at a time
		1<<1,
		1<<2,
		1<<3,
		1<<4,
		1<<5,
		1<<6,
		1<<7,
		1<<8,
		1<<9,
	}
	if len(calls) != len(expected) {
		t.Errorf("WalkCombinations(10, 1): expected %v, got %v.", expected, calls)
	}	else {	
		sort.Sort(sort.IntSlice(calls))
		for i, v := range(calls) {
			if v != expected[i] {
				t.Errorf("WalkCombinations(10, 1), call %d: expected %v, got %v.", i, expected[i], v)
			}
		}
	}	
}

func TestChooseAll(t *testing.T) {
	calls := make([]int, 0, 10)
	WalkCombinations(10, 10, func(perm []bool) {
		calls = append(calls, permToInt(perm))
	})
	expected := []int{
		1<<10 - 1,  // all should be true
	}
	if len(calls) != len(expected) {
		t.Errorf("Expected %v, got %v.", expected, calls)
	}	else {	
		for i, v := range(calls) {
			if v != expected[i] {
				t.Errorf("WalkCombinations(10, 10), call %d: expected %v, got %v.", i, expected[i], v)
			}
		}
	}	
}

func TestChooseSome(t *testing.T) {
	var calls uint = 0
	WalkCombinations(10, 4, func(perm []bool) {
		calls += 1
	})
	expected := Combinations(10, 4)
	if calls != expected {
		t.Errorf("WalkCombinations(10, 4): expected %v callbacks, got %v.", expected, calls)
	}
}
