package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Day7("input_example.txt")
	var expected uint64 = 95437
	if got != expected {
		t.Log("Got", got, " expected", expected)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	got := Day7part2("input_example.txt")
	var expected uint64 = 24933642
	if got != expected {
		t.Log("Got", got, " expected", expected)
		t.Fail()
	}
}
