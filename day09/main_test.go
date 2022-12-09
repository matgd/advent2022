package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Day9("input_example.txt")
	var expected uint64 = 13
	if got != expected {
		t.Log("Got", got, " expected", expected)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	// got := Day9_2("input_example.txt")
	// var expected uint64 = 8
	// if got != expected {
	// t.Log("Got", got, " expected", expected)
	// t.Fail()
	// }
}
