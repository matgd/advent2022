package main

import (
	"testing"
)

func TestTask1(t *testing.T) {
	got := Task1("input_example.txt")
	var expected uint32 = 15
	if got != expected {
		t.Log("Got", got, " expected", expected)
		t.Fail()
	}
}

func TestTask2(t *testing.T) {
	got := Task2("input_example.txt")
	var expected uint32 = 12
	if got != expected {
		t.Log("Got", got, " expected", expected)
		t.Fail()
	}
}
