package day2

import "testing"

func Test_level1(t *testing.T) {
	if level1() != 2512 {
		t.Log("wrong result")
		t.Fail()
	}
}

func Test_level2(t *testing.T) {
	if level2() != 67335 {
		t.Log("wrong result")
		t.Fail()
	}
}
