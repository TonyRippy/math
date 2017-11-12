package math

import (
	"sort"
	"testing"
)

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
	GenerateCombinations(0, 0, func(perm []bool) {
		calls += 1
		if len(perm) != 0 {
			t.Errorf("Expected empty permutation, got %v", perm)
		}
	})
	if calls != 1 {
		t.Errorf("Expected 1 callback from Choose(0,0), got %d", calls)
	}
	calls = 0
	GenerateCombinations(100, 0, func(perm []bool) {
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
		t.Errorf("Expected 1 callback from Choose(0,100), got %d", calls)
	}	
}

func TestChooseOne(t *testing.T) {
	calls := make([]int, 0, 10)
	GenerateCombinations(10, 1, func(perm []bool) {
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
		t.Errorf("Expected %v, got %v.", expected, calls)
	}	else {	
		sort.Sort(sort.IntSlice(calls))
		for i, v := range(calls) {
			if v != expected[i] {
				t.Errorf("Call %d: expected %v, got %v.", i, expected[i], v)
			}
		}
	}	
}

func TestChooseAll(t *testing.T) {
	calls := make([]int, 0, 10)
	GenerateCombinations(10, 10, func(perm []bool) {
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
				t.Errorf("Call %d: expected %v, got %v.", i, expected[i], v)
			}
		}
	}	
}

func TestChooseSome(t *testing.T) {
	calls := 0
	GenerateCombinations(10, 4, func(perm []bool) {
		calls += 1
	})
	// 10 choose 4 = (10! / (4! * 6!))
	expected := 210
	if calls != expected {
		t.Errorf("Expected %v callbacks, got %v.", expected, calls)
	}
}
