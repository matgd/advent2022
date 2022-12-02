package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Day11("input_example.txt")
	var expected int = 10605
	if got != expected {
		t.Log("Got", got, " expected", expected)
		t.Fail()
	}
}
