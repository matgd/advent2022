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
	ropeTail    *Rope
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

func (r Rope) RopeTailNeedsReadjustment() bool {
	if math.Abs(float64(r.ropeTail.head.X-r.head.X)) > 1 {
		return true
	}
	if math.Abs(float64(r.ropeTail.head.Y-r.head.Y)) > 1 {
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

func (r *Rope) ReadjustRopeTail() {
	if !r.RopeTailNeedsReadjustment() {
		return
	}

	xDiff := r.head.X - r.ropeTail.head.X
	yDiff := r.head.Y - r.ropeTail.head.Y

	if xDiff > 0 {
		r.ropeTail.head.X++
	}
	if xDiff < 0 {
		r.ropeTail.head.X--
	}
	if yDiff > 0 {
		r.ropeTail.head.Y++
	}
	if yDiff < 0 {
		r.ropeTail.head.Y--
	}

	r.tailVisited[Cord{X: r.ropeTail.head.X, Y: r.ropeTail.head.Y}] = nil
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

func Day9p2(filename string) uint64 {
	lines := utils.ReadFileLines(filename)

	ropes := make([]*Rope, 10)

	root := Rope{
		head:        Cord{X: 0, Y: 0},
		tailVisited: make(map[Cord]interface{}),
		ropeTail:    nil,
	}
	ropes[0] = &root

	for i := 1; i < 10; i++ {
		ropes[i] = &Rope{
			head:        Cord{X: 0, Y: 0},
			tailVisited: make(map[Cord]interface{}),
			ropeTail:    nil,
		}
		ropes[i-1].ropeTail = ropes[i]
	}

	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		direction := splitLine[0]
		distance, _ := strconv.Atoi(splitLine[1])

		for i := 0; i < distance; i++ {
			ropes[0].MoveHead(direction)
			ropes[0].ReadjustRopeTail()
			for i := 1; i < 10; i++ {
				if ropes[i].ropeTail != nil {
					ropes[i].ReadjustRopeTail()
				}
			}
		}
	}

	return uint64(len(ropes[8].tailVisited)) + 1 // +1 for the starting point
}

func main() {
	fmt.Println("Day 9, Part 1:", Day9("input.txt"))
	fmt.Println("Day 9, Part 2:", Day9p2("input.txt"))
}
