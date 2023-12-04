package day4

import "testing"

func check(t *testing.T, result int, expected int) {
	if result != expected {
		t.Logf("wrong result!\nactual:   %v\nexpected: %v", result, expected)
		t.Fail()
	}
}

func Test_level1_example(t *testing.T) {
	check(t, level1("example.txt"), 13)
}

func Test_level1_input(t *testing.T) {
	check(t, level1("input.txt"), 21213)
}

func Test_level2_example(t *testing.T) {
	check(t, level2("example.txt"), 30)
}

func Test_level2_input(t *testing.T) {
	check(t, level2("input.txt"), -1)
}
