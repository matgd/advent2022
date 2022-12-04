package main

import (
	"testing"
)

func TestTask1(t *testing.T) {
	got := Task1("input_example.txt")
	expected := 157
	if got != expected {
		t.Log("Got", got, " expected", expected)
		t.Fail()
	}
}

func TestTask2(t *testing.T) {
	got := Task2("input_example.txt")
	expected := 70
	if got != expected {
		t.Log("Got", got, " expected", expected)
		t.Fail()
	}
}
