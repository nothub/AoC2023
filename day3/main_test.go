package day3

import "testing"

func check(t *testing.T, result int, expected int) {
	if result != expected {
		t.Logf("wrong result!\nactual:   %v\nexpected: %v", result, expected)
		t.Fail()
	}
}

func Test_level1_example(t *testing.T) {
	check(t, level1("example.txt"), 4361)
}

func Test_level1_example2(t *testing.T) {
	check(t, level1("example2.txt"), 4462)
}

func Test_level1_input(t *testing.T) {
	check(t, level1("input.txt"), -1)
}
