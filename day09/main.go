package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/matgd/advent2022/utils"
)

type Cord struct {
	X, Y int
}

type Rope struct {
	head        Cord
	tail        Cord
	tailVisited map[Cord]interface{}
}

func (r Rope) TailNeedsReadjustment() bool {
	if math.Abs(float64(r.tail.X-r.head.X)) > 1 {
		return true
	}
	if math.Abs(float64(r.tail.Y-r.head.Y)) > 1 {
		return true
	}

	return false
}

func (r *Rope) ReadjustTail() {
	if !r.TailNeedsReadjustment() {
		return
	}

	xDiff := r.head.X - r.tail.X
	yDiff := r.head.Y - r.tail.Y

	if xDiff > 0 {
		r.tail.X++
	}
	if xDiff < 0 {
		r.tail.X--
	}
	if yDiff > 0 {
		r.tail.Y++
	}
	if yDiff < 0 {
		r.tail.Y--
	}

	r.tailVisited[Cord{X: r.tail.X, Y: r.tail.Y}] = nil
}

func (r *Rope) MoveHead(direction string) {
	switch direction {
	case "U":
		r.head.Y++
	case "D":
		r.head.Y--
	case "L":
		r.head.X--
	case "R":
		r.head.X++
	}
}

func Day9(filename string) uint64 {
	lines := utils.ReadFileLines(filename)

	r := Rope{
		head:        Cord{X: 0, Y: 0},
		tail:        Cord{X: 0, Y: 0},
		tailVisited: make(map[Cord]interface{}),
	}

	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		direction := splitLine[0]
		distance, _ := strconv.Atoi(splitLine[1])

		for i := 0; i < distance; i++ {
			r.MoveHead(direction)
			r.ReadjustTail()
		}
	}
	return uint64(len(r.tailVisited)) + 1 // +1 for the starting point
}

func main() {
	fmt.Println("Day 9, Part 1:", Day9("input.txt"))
}
