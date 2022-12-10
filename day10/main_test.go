package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Day10("input_example.txt")
	var expected int = 13140
	if got != expected {
		t.Log("Got", got, " expected", expected)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	// got := Day10_2("input_example.txt")
	// var expected uint64 = 8
	// if got != expected {
	// t.Log("Got", got, " expected", expected)
	// t.Fail()
	// }
}
