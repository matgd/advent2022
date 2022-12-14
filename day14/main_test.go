package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Day14("input_example.txt", 1)
	var expected int = 24
	if got != expected {
		t.Log("Got", got, " expected", expected)
		t.Fail()
	}
}
