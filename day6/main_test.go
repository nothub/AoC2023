package day6

import "testing"

func check(t *testing.T, result int, expected int) {
	if result != expected {
		t.Logf("wrong result!\nactual:   %v\nexpected: %v", result, expected)
		t.Fail()
	}
}

func Test_level1_input(t *testing.T) {
	check(t, level1("input.txt"), 32076)
}
