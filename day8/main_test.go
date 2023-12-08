package day8

import "testing"

func check(t *testing.T, result int, expected int) {
	if result != expected {
		t.Logf("wrong result!\nactual:   %v\nexpected: %v", result, expected)
		t.Fail()
	}
}

func Test_level1_example1(t *testing.T) {
	check(t, level1("example1.txt"), 2)
}

func Test_level1_example2(t *testing.T) {
	check(t, level1("example2.txt"), 6)
}

func Test_level1_input(t *testing.T) {
	check(t, level1("input.txt"), 18673)
}

func Test_level2_example(t *testing.T) {
	check(t, level2("example.txt"), -1)
}

func Test_level2_input(t *testing.T) {
	check(t, level2("input.txt"), -1)
}
