package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Day8("input_example.txt")
	var expected uint64 = 21
	if got != expected {
		t.Log("Got", got, " expected", expected)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	got := Day8_2("input_example.txt")
	var expected uint64 = 8
	if got != expected {
		t.Log("Got", got, " expected", expected)
		t.Fail()
	}
}
