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
	got := Day9p2("input_example_bigger.txt")
	var expected uint64 = 36
	if got != expected {
		t.Log("Got", got, " expected", expected)
		t.Fail()
	}
}
